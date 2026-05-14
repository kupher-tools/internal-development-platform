package domain

import (
	"time"

	"github.com/google/uuid"
)

type CreateContainerRepositoryRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	IsPrivate   bool   `json:"isPrivate"`
}

type ContainerRepository struct {
	ID          uuid.UUID
	Name        string
	Description string
	Provider    string
	Visibility  string

	DesiredState string
	ActualState  string
	Status       string

	Generation         int
	ObservedGeneration int

	CreatedAt time.Time
}

type RepoResponse struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
}
