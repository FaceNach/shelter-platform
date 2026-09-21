package handler

import (
	"net/http"
	"time"
)

type createShelterRequest struct {
	Name        string     `json:"name" validate:"required,max=100"`
	City        string     `json:"city" validate:"required"`
	State       string     `json:"state" validate:"required"`
	Country     string     `json:"country" validate:"required"`
	ZIPCode     string     `json:"zip_code" validate:"required"`
	Email       string     `json:"email" validate:"required,email"`
	PhoneNumber string     `json:"phone_number" validate:"required"`
	FoundedAt   *time.Time `json:"founded_at" validate:"required"`
}

func (h *Handler) createShelter(w http.ResponseWriter, r *http.Request) {

}
