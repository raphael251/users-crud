package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/raphael251/users-crud/internal/domain/user"
	"github.com/raphael251/users-crud/internal/domain/utils"
	responsehelpers "github.com/raphael251/users-crud/internal/infra/web/response_helpers"
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
		responsehelpers.BadRequest(w, r, []string{"Invalid JSON. Please see the docs."})
		return
	}

	if validationErrors := receivedUser.Validate(); validationErrors != nil {
		errs := make([]string, 0)
		for _, err := range validationErrors {
			errs = append(errs, err.Error())
		}

		responsehelpers.BadRequest(w, r, errs)
		return
	}

	result, err := h.CreateUserUseCase.Execute(receivedUser)

	if err != nil {
		if cerr, ok := err.(*utils.UseCaseError); ok {
			if cerr.Type == utils.Validation {
				responsehelpers.BadRequest(w, r, []string{cerr.Error()})
			}

			if cerr.Type == utils.BusinessRuleViolation {
				responsehelpers.BadRequest(w, r, []string{cerr.Error()})
			}
		}

		log.Println(err.Error())
		responsehelpers.InternalServerError(w, r)
		return
	}

	json.NewEncoder(w).Encode(result)
}
