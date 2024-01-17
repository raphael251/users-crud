package routers

import (
	"database/sql"

	"github.com/go-chi/chi/v5"
	"github.com/raphael251/users-crud/internal/domain/user"
	"github.com/raphael251/users-crud/internal/infra/database"
	"github.com/raphael251/users-crud/internal/infra/web/handlers"
)

func GenerateUserRouter(db *sql.DB) func(r chi.Router) {
	userRepository := database.NewUserRepository(db)
	createUserUseCase := user.NewCreateUserUseCase(userRepository)
	updateUserUseCase := user.NewUpdateUserUseCase(userRepository)
	userHandler := handlers.NewUserHandler(userRepository, createUserUseCase, updateUserUseCase)

	return func(r chi.Router) {
		r.Post("/", userHandler.Create)
		r.Put("/{id}", userHandler.Update)
	}
}
