package routers

import (
	"github.com/go-chi/chi/v5"
	"github.com/raphael251/users-crud/internal/infra/web/handlers"
)

func GenerateApplicationRouter() func(r chi.Router) {
	applicationHandler := handlers.NewApplicationHandler()

	return func(r chi.Router) {
		r.Get("/", applicationHandler.Health)
	}
}
