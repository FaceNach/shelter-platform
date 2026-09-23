package service

import (
	"context"
	"time"

	"shelter-platform/internal/shelter/domain"
)

type CreateShelterInput struct {
	Name         string
	City         string
	State        string
	Country      string
	ZIPCode      string
	ContactEmail string
	PhoneNumber  string
	FoundedAt    *time.Time
}

func (s *Service) CreateShelter(ctx context.Context, input CreateShelterInput) (domain.Shelter, error) {

	shelter, err := domain.NewShelter(domain.NewShelterParams{
		Name:         input.Name,
		City:         input.City,
		State:        input.State,
		Country:      input.Country,
		ZIPCode:      input.ZIPCode,
		PhoneNumber:  input.PhoneNumber,
		ContactEmail: input.ContactEmail,
		FoundedAt:    input.FoundedAt,
	})

	if err != nil {
		return domain.Shelter{}, err
	}

	return s.repo.Create(ctx, shelter)
}
