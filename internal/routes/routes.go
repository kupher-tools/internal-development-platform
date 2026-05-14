package routes

import (
	"internal-development-platform/internal/config"
	"internal-development-platform/internal/handler"
	"internal-development-platform/internal/middleware"
	"internal-development-platform/internal/persistence"
	"internal-development-platform/internal/service"
	"net/http"

	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"
)

func SetupRoutes(cfg *config.Config, postgresClient *persistence.Postgres) http.Handler {

	r := chi.NewRouter()
	r.Use(middleware.Logging)

	containerRepoStore :=
		persistence.NewContainerRepositoryStore(
			postgresClient.DB,
		)

	healthSvc := service.NewHealthService("3.1.2")
	containerRepoSvc := service.NewContainerRepoService(cfg.ContainerRepo, containerRepoStore)

	healthHandler := handler.NewHealthHandler(healthSvc)
	containerRepoHandler := handler.NewContainerRepoHandler(containerRepoSvc)

	r.Get("/health/live", healthHandler.Live)
	r.Get("/health/ready", healthHandler.Ready)
	r.Get("/swagger/*", httpSwagger.WrapHandler)

	r.Route("/api/v1", func(api chi.Router) {

		api.Route("/container-repositories", func(repo chi.Router) {
			repo.Post("/", containerRepoHandler.Create)

			// Future
			// repo.Get("/", containerRepoHandler.List)
			// repo.Get("/{id}", containerRepoHandler.Get)
			// repo.Delete("/{id}", containerRepoHandler.Delete)
		})

		// Future APIs

		// api.Route("/source-repositories", func(r chi.Router) {})
		// api.Route("/code-quality-projects", func(r chi.Router) {})
		// api.Route("/applications", func(r chi.Router) {})
	})

	return r
}
