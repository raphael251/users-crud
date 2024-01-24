package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/raphael251/users-crud/internal/domain/usecase"
	"github.com/raphael251/users-crud/internal/domain/utils"
	"github.com/raphael251/users-crud/internal/infra/web/httputils"
)

func CreateUserHandler(useCase *usecase.CreateUserUseCase) func(w http.ResponseWriter, r *http.Request) {
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

func UpdateUserHandler(useCase *usecase.UpdateUserUseCase) func(w http.ResponseWriter, r *http.Request) {
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

func FindOneUserHandler(useCase *usecase.FindOneUserUseCase) func(w http.ResponseWriter, r *http.Request) {
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
