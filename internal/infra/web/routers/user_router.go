package routers

import (
	"database/sql"

	"github.com/go-chi/chi/v5"
	"github.com/raphael251/users-crud/internal/infra/database"
	"github.com/raphael251/users-crud/internal/infra/web/handlers"
)

func GenerateUserRouter(db *sql.DB) func(r chi.Router) {
	userRepository := database.NewUserRepository(db)

	return func(r chi.Router) {
		r.Post("/", handlers.CreateUserHandler(userRepository))
		r.Get("/{id}", handlers.FindOneUserHandler(userRepository))
		r.Put("/{id}", handlers.UpdateUserHandler(userRepository))
	}
}
