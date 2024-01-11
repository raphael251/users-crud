package handlers

import (
	"net/http"

	"github.com/raphael251/users-crud/internal/user"
)

type UserHandler struct {
	UserRepository user.UserRepositoryInterface
}

func NewUserHandler(userRepository user.UserRepositoryInterface) *UserHandler {
	return &UserHandler{
		UserRepository: userRepository,
	}
}

func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	// controller
	// parse the r.Body
}
