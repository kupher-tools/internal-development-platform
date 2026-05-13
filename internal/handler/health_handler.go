package handler

import (
	"net/http"

	appErrors "internal-development-platform/internal/errors"
	"internal-development-platform/internal/ports"
	"internal-development-platform/internal/response"
)

type HealthHandler struct {
	svc ports.HealthService
}

func NewHealthHandler(svc ports.HealthService) *HealthHandler {
	return &HealthHandler{svc: svc}
}

// Live checks the liveness of the application
// @Summary      Get Liveness
// @Description  Returns the health status of the instance
// @Tags         health
// @Produce      json
// @Success 200 {object} map[string]string
// @Failure 503 {object} map[string]string
// @Router       /health/live [get]
func (h *HealthHandler) Live(w http.ResponseWriter, r *http.Request) {
	err := h.svc.CheckLiveness(r.Context())
	if err != nil {
		appErrors.HandleHTTPError(w, err)
		return

	}
	response.HandleHTTPResponse(w, http.StatusOK, map[string]string{
		"status": "UP",
	})

}

// Ready checks the readiness of the application
// @Summary      Get Readiness
// @Description  Returns the health status of the instance
// @Tags         health
// @Produce      json
// @Success 200 {object} map[string]string
// @Failure 503 {object} map[string]string
// @Router       /health/ready [get]
func (h *HealthHandler) Ready(w http.ResponseWriter, r *http.Request) {
	err := h.svc.CheckReadiness(r.Context())
	if err != nil {
		appErrors.HandleHTTPError(w, err)
		return

	}
	response.HandleHTTPResponse(w, http.StatusOK, map[string]string{
		"status": "UP",
	})
}
