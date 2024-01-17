package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/raphael251/users-crud/internal/domain/interfaces"
	"github.com/raphael251/users-crud/internal/domain/usecase"
	"github.com/raphael251/users-crud/internal/domain/utils"
	"github.com/raphael251/users-crud/internal/infra/web/httputils"
)

type UserHandler struct {
	UserRepository     interfaces.UserRepositoryInterface
	CreateUserUseCase  usecase.CreateUserUseCase
	UpdateUserUseCase  usecase.UpdateUserUseCase
	FindOneUserUseCase usecase.FindOneUserUseCase
}

func NewUserHandler(
	userRepository interfaces.UserRepositoryInterface,
	createUserUseCase *usecase.CreateUserUseCase,
	updateUserUseCase *usecase.UpdateUserUseCase,
	findOneUserUseCase *usecase.FindOneUserUseCase,
) *UserHandler {
	return &UserHandler{
		UserRepository:     userRepository,
		CreateUserUseCase:  *createUserUseCase,
		UpdateUserUseCase:  *updateUserUseCase,
		FindOneUserUseCase: *findOneUserUseCase,
	}
}

func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var receivedUser usecase.CreateUserInputDTO

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

	contentLocation := fmt.Sprintf("/api/v1/users/%s", result.ID)
	httputils.RespondCreated(w, r, contentLocation)
}

func (h *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	var receivedUser usecase.UpdateUserInputDTO

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

	receivedUser.ID = chi.URLParam(r, "id")

	err = h.UpdateUserUseCase.Execute(receivedUser)

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

	httputils.RespondNoContent(w, r)
}

func (h *UserHandler) FindOne(w http.ResponseWriter, r *http.Request) {
	receivedUserId := chi.URLParam(r, "id")

	user, err := h.FindOneUserUseCase.Execute(receivedUserId)

	if err != nil {
		if cerr, ok := err.(*utils.UseCaseError); ok {
			if cerr.Type == utils.ValidationError {
				httputils.RespondBadRequest(w, r, []string{cerr.Error()})
				return
			}

			if cerr.Type == utils.NotFoundError {
				httputils.RespondNotFound(w, r)
				return
			}
		}

		log.Println(err.Error())
		httputils.RespondInternalServerError(w, r)
		return
	}

	httputils.RespondOK(w, r, user)
}
