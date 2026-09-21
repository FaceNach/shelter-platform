package domain

import "context"

type Repository interface {
	Create(ctx context.Context, shelter Shelter) (Shelter, error)
}
