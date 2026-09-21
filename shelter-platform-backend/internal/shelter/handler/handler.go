package handler

import (
	"fmt"
	"net/http"
	"shelter-platform/internal/shelter/service"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

type Handler struct {
	service   *service.Service
	validator *validator.Validate
}

func New(service *service.Service, validator *validator.Validate) *Handler {
	return &Handler{
		service:   service,
		validator: validator,
	}
}

func (h *Handler) Routes() http.Handler {
	r := chi.NewRouter()

	r.Get("/hello", h.helloWorld)
	r.Post("/createShelter", h.createShelter)

	return r
}

func (h *Handler) helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World!")
}
