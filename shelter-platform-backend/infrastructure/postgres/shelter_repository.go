package postgres

import (
	"context"

	"shelter-platform/infrastructure/postgres/sqlcgen"
	"shelter-platform/internal/shelter/domain"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ShelterRepository struct {
	queries *sqlcgen.Queries
}

func NewShelterRepository(pool *pgxpool.Pool) *ShelterRepository {
	return &ShelterRepository{
		queries: sqlcgen.New(pool),
	}
}

func (r *ShelterRepository) Create(ctx context.Context, shelter domain.Shelter) (domain.Shelter, error) {
	foundedAt := pgtype.Date{}
	if shelter.FoundedAt != nil {
		foundedAt.Time = *shelter.FoundedAt
		foundedAt.Valid = true
	}

	row, err := r.queries.CreateShelter(ctx, sqlcgen.CreateShelterParams{
		Name:         shelter.Name,
		City:         shelter.City,
		State:        shelter.State,
		Country:      shelter.Country,
		ZipCode:      shelter.ZIPCode,
		PhoneNumber:  shelter.PhoneNumber,
		ContactEmail: shelter.ContactEmail,
		FoundedAt:    foundedAt,
	})
	if err != nil {
		return domain.Shelter{}, err
	}

	created := domain.Shelter{
		ID:           row.ID,
		Name:         row.Name,
		City:         row.City,
		State:        row.State,
		Country:      row.Country,
		ZIPCode:      row.ZipCode,
		PhoneNumber:  row.PhoneNumber,
		ContactEmail: row.ContactEmail,
		CreatedAt:    row.CreatedAt.Time,
		UpdatedAt:    row.UpdatedAt.Time,
	}
	if row.FoundedAt.Valid {
		created.FoundedAt = &row.FoundedAt.Time
	}

	return created, nil
}
