package domain

import (
	"fmt"
	"strings"
	"time"
)

type Shelter struct {
	ID           int64
	Name         string
	City         string
	State        string
	Country      string
	ZIPCode      string
	PhoneNumber  string
	ContactEmail string
	FoundedAt    *time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type NewShelterParams struct {
	Name         string
	City         string
	State        string
	Country      string
	ZIPCode      string
	PhoneNumber  string
	ContactEmail string
	FoundedAt    *time.Time
}

func NewShelter(params NewShelterParams) (Shelter, error) {
	name := strings.TrimSpace(params.Name)
	city := strings.TrimSpace(params.City)
	state := strings.TrimSpace(params.State)
	country := strings.TrimSpace(params.Country)
	zipCode := strings.TrimSpace(params.ZIPCode)
	phoneNumber := strings.TrimSpace(params.PhoneNumber)
	email := strings.TrimSpace(params.ContactEmail)

	if name == "" {
		return Shelter{}, fmt.Errorf("name cannot be empty")
	}

	if len(name) <= 3 {
		return Shelter{}, fmt.Errorf("name needs to have more than 3 characters")
	}

	if city == "" {
		return Shelter{}, fmt.Errorf("city cannot be empty")
	}

	if state == "" {
		return Shelter{}, fmt.Errorf("state cannot be empty")
	}

	if country == "" {
		return Shelter{}, fmt.Errorf("country cannot be empty")
	}

	if zipCode == "" {
		return Shelter{}, fmt.Errorf("zip code cannot be empty")
	}

	if phoneNumber == "" {
		return Shelter{}, fmt.Errorf("phone number cannot be empty")
	}

	if email == "" {
		return Shelter{}, fmt.Errorf("email cannot be empty")
	}

	if params.FoundedAt == nil {
		return Shelter{}, fmt.Errorf("founded at cannot be nil")
	}

	if params.FoundedAt.After(time.Now()) {
		return Shelter{}, fmt.Errorf("founded at cannot be in the future")
	}

	return Shelter{
		Name:         name,
		City:         city,
		State:        state,
		Country:      country,
		ZIPCode:      zipCode,
		PhoneNumber:  phoneNumber,
		ContactEmail: email,
		FoundedAt:    params.FoundedAt,
	}, nil
}
