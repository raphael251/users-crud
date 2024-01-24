package routers

import (
	"database/sql"

	"github.com/go-chi/chi/v5"
	"github.com/raphael251/users-crud/internal/domain/usecase"
	"github.com/raphael251/users-crud/internal/infra/database"
	"github.com/raphael251/users-crud/internal/infra/web/handlers"
)

func GenerateUserRouter(db *sql.DB) func(r chi.Router) {
	userRepository := database.NewUserRepository(db)
	createUserUseCase := usecase.NewCreateUserUseCase(userRepository)
	updateUserUseCase := usecase.NewUpdateUserUseCase(userRepository)
	findOneUserUseCase := usecase.NewFindOneUserUseCase(userRepository)

	return func(r chi.Router) {
		r.Post("/", handlers.CreateUserHandler(createUserUseCase))
		r.Get("/{id}", handlers.FindOneUserHandler(findOneUserUseCase))
		r.Put("/{id}", handlers.UpdateUserHandler(updateUserUseCase))
	}
}
