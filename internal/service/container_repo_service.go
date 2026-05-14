package service

import (
	"context"
	"internal-development-platform/internal/config"
	"internal-development-platform/internal/domain"
	appErrors "internal-development-platform/internal/errors"
	"internal-development-platform/internal/ports"
	"log/slog"

	"github.com/google/uuid"
)

type ContainerRepoService struct {
	containerRepo config.ContainerRepo
	store         ports.ContainerRepositoryStore
}

func NewContainerRepoService(containerRepo config.ContainerRepo, store ports.ContainerRepositoryStore) *ContainerRepoService {

	return &ContainerRepoService{
		containerRepo: containerRepo,
		store:         store,
	}
}

func (s *ContainerRepoService) CreateRepository(
	ctx context.Context,
	req domain.CreateContainerRepositoryRequest,
) (*domain.RepoResponse, error) {

	if req.Name == "" {
		return nil, appErrors.ErrInternal
	}

	visibility := "public"

	if req.IsPrivate {
		visibility = "private"
	}

	repo := domain.ContainerRepository{
		ID:          uuid.New(),
		Name:        req.Name,
		Description: req.Description,
		Provider:    "harbor",
		Visibility:  visibility,

		DesiredState: "present",
		ActualState:  "pending",
		Status:       "pending",

		Generation:         1,
		ObservedGeneration: 0,
	}

	err := s.store.Create(ctx, repo)
	if err != nil {
		slog.Error("Err", "err", err)
		return nil, appErrors.ErrInternal
	}

	return &domain.RepoResponse{
		ID:     repo.ID.String(),
		Name:   repo.Name,
		Status: repo.Status,
	}, nil
}

/*
func (s *ContainerRepoService) CreateRepository(ctx context.Context, req domain.CreateContainerRepositoryRequest) (*domain.RepoResponse, error) {
	if req.Name == "" {
		return nil, appErrors.ErrInternal
	}

	visibility := "public"
	if req.IsPrivate {
		visibility = "private"
	}
	repo := domain.ContainerRepository{
		ID:                 uuid.New(),
		Name:               req.Name,
		Description:        req.Description,
		Provider:           "harbor",
		Visibility:         visibility,
		DesiredState:       "present",
		Status:             "pending",
		Generation:         1,
		ObservedGeneration: 0,
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
*/
