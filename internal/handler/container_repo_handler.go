package handler

import (
	"encoding/json"
	"internal-development-platform/internal/domain"
	appErrors "internal-development-platform/internal/errors"
	"internal-development-platform/internal/ports"
	"internal-development-platform/internal/response"
	"net/http"
)

type ContainerRepoHandler struct {
	svc ports.ContainerRepoService
}

func NewContainerRepoHandler(svc ports.ContainerRepoService) *ContainerRepoHandler {
	return &ContainerRepoHandler{svc: svc}
}

// Create handles the creation of a new repository
// @Summary      Create Container Repository
// @Description  Creates a new container repository
// @Tags         repository
// @Accept       json
// @Produce      json
// @Param        request body domain.RepoRequest true "Create container repository"
// @Success      201  {object}  domain.RepoResponse
// @Router       /api/v1/container-repositories [post]
func (h *ContainerRepoHandler) Create(w http.ResponseWriter, r *http.Request) {
	var repoRequest domain.CreateContainerRepositoryRequest

	if err := json.NewDecoder(r.Body).Decode(&repoRequest); err != nil {
		appErrors.HandleHTTPError(w, err)
		return
	}

	repo, err := h.svc.CreateRepository(r.Context(), repoRequest)
	if err != nil {
		appErrors.HandleHTTPError(w, err)
		return
	}

	response.HandleHTTPResponse(w, http.StatusCreated, repo)

}
