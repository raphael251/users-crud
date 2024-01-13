package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/raphael251/users-crud/internal/user"
	"github.com/raphael251/users-crud/pkg/utils"
)

type UserHandler struct {
	UserRepository    user.UserRepositoryInterface
	CreateUserUseCase user.CreateUserUseCase
}

func NewUserHandler(userRepository user.UserRepositoryInterface, createUserUseCase *user.CreateUserUseCase) *UserHandler {
	return &UserHandler{
		UserRepository:    userRepository,
		CreateUserUseCase: *createUserUseCase,
	}
}

func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var receivedUser user.CreateUserInputDTO

	err := json.NewDecoder(r.Body).Decode(&receivedUser)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)

		errMessage := utils.Error{Message: "bad request", Data: []string{"Invalid JSON. Please see the docs."}}
		json.NewEncoder(w).Encode(errMessage)
		return
	}

	if validationErrors := receivedUser.Validate(); validationErrors != nil {
		errs := make([]string, 0)
		for _, err := range validationErrors {
			errs = append(errs, err.Error())
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)

		errMessage := utils.Error{Message: "bad request", Data: errs}
		json.NewEncoder(w).Encode(errMessage)
		return
	}

	result, err := h.CreateUserUseCase.Execute(receivedUser)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)

		errMessage := utils.Error{Message: "internal server error"}
		json.NewEncoder(w).Encode(errMessage)
		return
	}

	json.NewEncoder(w).Encode(result)
}
