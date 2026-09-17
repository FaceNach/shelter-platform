package app

import (
	"context"
	"fmt"
	"net/http"
	"shelter-platform/infrastructure/postgres"
	shelterHandler "shelter-platform/internal/shelter/handler"
	shelterService "shelter-platform/internal/shelter/service"
)

func Run(ctx context.Context) error {

	fmt.Println("Loading .env variables")
	cfg, err := LoadConfig()
	if err != nil {
		return err
	}

	fmt.Println("Starting connection to DB")
	pool, err := postgres.NewPool(ctx, cfg.databaseUrl)
	if err != nil {
		return err
	}
	defer pool.Close()

	shelterRepo := postgres.NewShelterRepository(pool)
	shelterService := shelterService.New(shelterRepo)
	shelterHandler := shelterHandler.New(shelterService)

	r := newRouter(shelterHandler)

	fmt.Printf("Server listening on port %s\n", cfg.port)
	return http.ListenAndServe(cfg.port, r)
}
