package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"shelter-platform/internal/httpapi"
	"shelter-platform/internal/shelter/domain"
	"shelter-platform/internal/shelter/service"

	"time"

	"github.com/go-playground/validator/v10"
)

type createShelterRequest struct {
	Name         string `json:"name" validate:"required,max=100"`
	City         string `json:"city" validate:"required"`
	State        string `json:"state" validate:"required"`
	Country      string `json:"country" validate:"required"`
	ZIPCode      string `json:"zip_code" validate:"required"`
	ContactEmail string `json:"contact_email" validate:"required,email"`
	PhoneNumber  string `json:"phone_number" validate:"required"`
	FoundedAt    string `json:"founded_at" validate:"required"`
}

type createShelterResponse struct {
	ID int64 `json:"id"`
}

func (h *Handler) createShelter(w http.ResponseWriter, r *http.Request) {

	var req createShelterRequest

	err := httpapi.DecodeJSONBody(w, r, &req)

	if err != nil {
		var mr *httpapi.MalformedRequest
		if errors.As(err, &mr) {
			http.Error(w, mr.Msg, mr.Status)
		} else {
			log.Print(err.Error())
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		return
	}

	err = h.validator.Struct(req)
	if err != nil {
		errors := err.(validator.ValidationErrors)
		http.Error(w, fmt.Sprintf("validation error: %v", errors), http.StatusBadRequest)
		return
	}

	foundedAt, err := time.Parse(time.RFC3339, req.FoundedAt)
	if err != nil {
		http.Error(w, "founded_at must be a valid RFC3339 timestamp", http.StatusBadRequest)
		return
	}

	shelterInput := service.CreateShelterInput{
		Name:         req.Name,
		City:         req.City,
		State:        req.State,
		Country:      req.Country,
		ZIPCode:      req.ZIPCode,
		ContactEmail: req.ContactEmail,
		PhoneNumber:  req.PhoneNumber,
		FoundedAt:    &foundedAt,
	}

	created, err := h.service.CreateShelter(r.Context(), shelterInput)
	if err != nil {
		var validationErr *domain.ValidationError
		if errors.As(err, &validationErr) {
			http.Error(w, validationErr.Error(), http.StatusBadRequest)
			return
		}

		log.Printf("create shelter: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	response := createShelterResponse{ID: created.ID}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("write create shelter response: %v", err)
	}
}
