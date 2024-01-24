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

func CreateUserHandler(repo interfaces.UserRepositoryInterface) func(w http.ResponseWriter, r *http.Request) {
	useCase := usecase.NewCreateUserUseCase(repo)

	return func(w http.ResponseWriter, r *http.Request) {
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

		result, err := useCase.Execute(receivedUser)

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
}

func UpdateUserHandler(repo interfaces.UserRepositoryInterface) func(w http.ResponseWriter, r *http.Request) {
	useCase := usecase.NewUpdateUserUseCase(repo)

	return func(w http.ResponseWriter, r *http.Request) {
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

		err = useCase.Execute(receivedUser)

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
}

func FindOneUserHandler(repo interfaces.UserRepositoryInterface) func(w http.ResponseWriter, r *http.Request) {
	useCase := usecase.NewFindOneUserUseCase(repo)

	return func(w http.ResponseWriter, r *http.Request) {
		receivedUserId := chi.URLParam(r, "id")

		user, err := useCase.Execute(receivedUserId)

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
}
