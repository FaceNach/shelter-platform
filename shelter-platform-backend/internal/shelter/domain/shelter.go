package domain

import (
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
		return Shelter{}, &ValidationError{
			Field:   "name",
			Message: "name cannot be empty",
		}
	}

	if len(name) <= 3 {
		return Shelter{}, &ValidationError{
			Field:   "name",
			Message: "name needs to have more than 3 characters",
		}
	}

	if city == "" {
		return Shelter{}, &ValidationError{
			Field:   "city",
			Message: "city cannot be empty",
		}
	}

	if state == "" {
		return Shelter{}, &ValidationError{
			Field:   "state",
			Message: "state cannot be empty",
		}
	}

	if country == "" {
		return Shelter{}, &ValidationError{
			Field:   "country",
			Message: "country cannot be empty",
		}
	}

	if zipCode == "" {
		return Shelter{}, &ValidationError{
			Field:   "zip_code",
			Message: "zip code cannot be empty",
		}
	}

	if phoneNumber == "" {
		return Shelter{}, &ValidationError{
			Field:   "phone_number",
			Message: "phone number cannot be empty",
		}
	}

	if email == "" {
		return Shelter{}, &ValidationError{
			Field:   "contact_email",
			Message: "email cannot be empty",
		}
	}

	if params.FoundedAt == nil {
		return Shelter{}, &ValidationError{
			Field:   "founded_at",
			Message: "founded at cannot be nil",
		}
	}

	if params.FoundedAt.After(time.Now()) {
		return Shelter{}, &ValidationError{
			Field:   "founded_at",
			Message: "founded at cannot be in the future",
		}
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
