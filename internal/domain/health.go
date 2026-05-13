package domain

type Status string

const (
	StatusUp   Status = "UP"
	StatusDown Status = "DOWN"
)

type Health struct {
	Status  Status            `json:"status"`
	Checks  map[string]Status `json:"checks,omitempty"`
	Version string            `json:"version,omitempty"`
}

func NewHealth(status Status, version string) Health {
	return Health{
		Status:  status,
		Version: version,
		Checks:  make(map[string]Status),
	}
}
