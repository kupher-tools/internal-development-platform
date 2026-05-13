package service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"internal-development-platform/internal/config"
	"internal-development-platform/internal/domain"
	appErrors "internal-development-platform/internal/errors"
	"net/http"
	"time"
)

type ContainerRepoService struct {
	containerRepo config.ContainerRepo
}

func NewContainerRepoService(containerRepo config.ContainerRepo) *ContainerRepoService {
	return &ContainerRepoService{
		containerRepo: containerRepo,
	}
}

func (s *ContainerRepoService) CreateRepository(ctx context.Context, repoRequest domain.RepoRequest) (*domain.RepoResponse, error) {
	if repoRequest.Name == "" {
		return nil, appErrors.ErrInternal
	}

	password := s.containerRepo.Password

	username := s.containerRepo.Username

	token, err := getDockerHubJWT(ctx, password, username)
	if err != nil {
		return nil, appErrors.ErrInternal
	}

	uri := fmt.Sprintf("%s/v2/namespaces/%s/repositories", s.containerRepo.URL, username)
	payload := repoRequest

	body, err := json.Marshal(payload)
	if err != nil {
		return nil, appErrors.ErrInternal
	}

	req, err := http.NewRequestWithContext(ctx, "POST", uri, bytes.NewBuffer(body))
	if err != nil {
		return nil, appErrors.ErrInternal
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "JWT "+token)

	client := &http.Client{Timeout: 15 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, appErrors.ErrInternal
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return nil, appErrors.ErrInternal
	}

	return domain.NewRepository(repoRequest.Name, uri), nil
}

func getDockerHubJWT(ctx context.Context, username, pat string) (string, error) {
	if username == "" || pat == "" {
		return "", appErrors.ErrInternal
	}

	payload := map[string]string{
		"username": username,
		"password": pat,
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return "", appErrors.ErrInternal
	}

	req, err := http.NewRequestWithContext(ctx, "POST", "https://hub.docker.com/v2/users/login/", bytes.NewBuffer(body))
	if err != nil {
		return "", appErrors.ErrInternal
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", appErrors.ErrInternal
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", appErrors.ErrInternal
	}

	var result struct {
		Token string `json:"token"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", appErrors.ErrInternal
	}
	if result.Token == "" {
		return "", appErrors.ErrInternal
	}

	return result.Token, nil
}
