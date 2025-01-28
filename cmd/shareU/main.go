package main

import (
	"github.com/labstack/echo/v4"
	"log/slog"
	"net/http"
	"os"
	"shareU/internal/config"
	v1 "shareU/internal/controller/http/v1"
	"shareU/internal/repo"
	"shareU/internal/service"
	"shareU/pkg/postgres"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	// Configuration
	cfg := config.MustLoad()

	// Logger
	log := setupLogger(cfg.Env)

	// Repositories
	log.Info("Initializing postgres...")
	pg, err := postgres.New(cfg.DBConfig.ConnectionURL(), postgres.MaxPoolSize(1))
	if err != nil {
		log.Error("app - Run - pgdb.NewServices: %w", err)
	}
	defer pg.Close()

	// Repositories
	log.Info("Initializing repositories...")
	repositories := repo.NewRepositories(pg, log)

	log.Info("Initializing services...")
	deps := service.ServicesDependencies{
		Repos: repositories,
	}
	services := service.NewServices(deps)

	// Echo handler
	log.Info("Initializing handlers and routes...")
	handler := echo.New()
	// setup handler validator as lib validator
	//handler.Validator = validator.NewCustomValidator()
	v1.NewRouter(handler, services)

	// HTTP server
	log.Info("Starting http server...")
	log.Debug("Server port: %s", cfg.Address)
	s := http.Server{
		Addr:    cfg.Address,
		Handler: handler,
	}
	if err := s.ListenAndServe(); err != http.ErrServerClosed {
		log.Error("error")
	}
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}
