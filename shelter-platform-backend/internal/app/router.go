package app

import (
	shelterHandler "shelter-platform/internal/shelter/handler"

	"github.com/go-chi/chi/v5"
)

func newRouter(shelterHandler *shelterHandler.Handler) chi.Router {
	r := chi.NewRouter()

	r.Mount("/v1/shelter", shelterHandler.Routes())

	return r
}
