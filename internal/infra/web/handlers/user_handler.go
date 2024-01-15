package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/raphael251/users-crud/internal/domain/user"
	"github.com/raphael251/users-crud/internal/domain/utils"
	"github.com/raphael251/users-crud/internal/infra/web/httputils"
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
		httputils.RespondBadRequest(w, r, []string{"Invalid JSON. Please see the docs."})
		return
	}

	if validationErrors := receivedUser.Validate(); validationErrors != nil {
		errs := make([]string, 0)
		for _, err := range validationErrors {
			errs = append(errs, err.Error())
		}

		httputils.RespondBadRequest(w, r, errs)
		return
	}

	result, err := h.CreateUserUseCase.Execute(receivedUser)

	if err != nil {
		if cerr, ok := err.(*utils.UseCaseError); ok {
			if cerr.Type == utils.ValidationError {
				httputils.RespondBadRequest(w, r, []string{cerr.Error()})
				return
			}

			if cerr.Type == utils.BusinessRuleViolationError {
				httputils.RespondBadRequest(w, r, []string{cerr.Error()})
				return
			}
		}

		log.Println(err.Error())
		httputils.RespondInternalServerError(w, r)
		return
	}

	json.NewEncoder(w).Encode(result)
}
