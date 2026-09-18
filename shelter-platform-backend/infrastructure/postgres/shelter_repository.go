package postgres

import "github.com/jackc/pgx/v5/pgxpool"

type ShelterRepository struct {
	//queries
	pool *pgxpool.Pool
}

func NewShelterRepository(pool *pgxpool.Pool) *ShelterRepository {
	return &ShelterRepository{
		pool: pool,
	}
}
