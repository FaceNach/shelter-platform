package app

import (
	"context"
	"net/http"
	"shelter-platform/infrastructure/postgres"
	shelterHandler "shelter-platform/internal/shelter/handler"
	shelterService "shelter-platform/internal/shelter/service"
)

func Run(ctx context.Context) error {
	cfg, err := LoadConfig()
	if err != nil {
		return err
	}

	pool, err := postgres.NewPool(ctx, cfg.databaseUrl)
	if err != nil {
		return err
	}
	defer pool.Close()

	shelterRepo := postgres.NewShelterRepository(pool)
	shelterService := shelterService.New(shelterRepo)
	shelterHandler := shelterHandler.New(shelterService)

	r := newRouter(shelterHandler)

	return http.ListenAndServe(cfg.port, r)
}
