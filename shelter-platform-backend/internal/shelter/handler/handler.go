package handler

import (
	"net/http"
	"shelter-platform/internal/shelter/service"

	"github.com/go-chi/chi/v5"
)

type Handler struct {
	service *service.Service
}

func New(service *service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) Routes() http.Handler {
	r := chi.NewRouter()

	return r
}
