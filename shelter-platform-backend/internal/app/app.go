package app

import (
	"context"
	"shelter-platform/infrastructure/postgres"
	"shelter-platform/internal/shelter/service"
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
	_ = service.New(shelterRepo)

	return nil
}
