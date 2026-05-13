package domain

type RepoRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	IsPrivate   bool   `json:"isPrivate"`
}

type RepoResponse struct {
	Name string `json:"name"`
	URI  string `json:"uri"`
	Type string `json:"type"` // e.g., "docker", "helm"
}

func NewRepository(name string, uri string) *RepoResponse {
	return &RepoResponse{
		Name: name,
		URI:  uri,
		Type: "docker",
	}
}
