package postgres

import "github.com/jackc/pgx/v5/pgxpool"

type ShelterRepository struct {
	pool *pgxpool.Pool
}

func NewShelterRepository(pool *pgxpool.Pool) *ShelterRepository {
	return &ShelterRepository{
		pool: pool,
	}
}
