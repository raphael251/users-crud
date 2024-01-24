package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/go-sql-driver/mysql"

	"github.com/raphael251/users-crud/configs"
	"github.com/raphael251/users-crud/internal/domain/usecase"
	"github.com/raphael251/users-crud/internal/domain/utils"
	"github.com/raphael251/users-crud/internal/infra/database"
	"github.com/raphael251/users-crud/internal/infra/web/httputils"
)

func main() {
	configs, err := configs.LoadConfig()
	if err != nil {
		panic(err)
	}

	db, err := sql.Open(configs.DBDriver, fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		configs.DBUser, configs.DBPassword, configs.DBHost, configs.DBPort, configs.DBName,
	))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Route("/api/v1/health", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(200)
		})
	})

	userRepository := database.NewUserRepository(db)

	router.Route("/api/v1/users", func(r chi.Router) {
		r.Post("/", func(w http.ResponseWriter, r *http.Request) {
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

			createUserUseCase := usecase.NewCreateUserUseCase(userRepository)
			result, err := createUserUseCase.Execute(receivedUser)

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
		})

		r.Get("/{id}", func(w http.ResponseWriter, r *http.Request) {
			receivedUserId := chi.URLParam(r, "id")

			findOneUserUseCase := usecase.NewFindOneUserUseCase(userRepository)
			user, err := findOneUserUseCase.Execute(receivedUserId)

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
		})

		r.Put("/{id}", func(w http.ResponseWriter, r *http.Request) {
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

			updateUserUseCase := usecase.NewUpdateUserUseCase(userRepository)
			err = updateUserUseCase.Execute(receivedUser)

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
		})
	})

	log.Printf("Server listening on port %s", configs.ServerPort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", configs.ServerPort), router))
}
