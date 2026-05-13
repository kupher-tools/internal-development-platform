package main

import (
	"log"
	"log/slog"
	"net/http"

	"internal-development-platform/internal/routes"

	_ "internal-development-platform/docs"

	"internal-development-platform/internal/config"
	"internal-development-platform/internal/logger"
)

// Version is injected at build time or loaded from config.
// Use: go build -ldflags "-X main.Version=1.0.0"
var Version = "dev"

// @title           Internal Development Platform API
// @version         1.0
// @description     Platform for managing internal engineering resources.
// @host            localhost:8080
// @BasePath        /
func main() {

	slog.SetDefault(logger.New())

	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	server := &http.Server{
		Addr:    cfg.Server.Port,
		Handler: routes.SetupRoutes(cfg),
	}
	slog.Info("Starting server", "port", cfg.Server.Port)

	if err := server.ListenAndServe(); err != nil {
		slog.Error("Server failed", "error", err)
	}
}
