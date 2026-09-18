package domain

import "time"

type Shelter struct {
	ID          int64
	Name        string
	City        string
	State       string
	Country     string
	ZIPCode     string
	PhoneNumber string
	FoundedAt   *time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
