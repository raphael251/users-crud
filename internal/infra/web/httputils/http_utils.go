package httputils

import (
	"encoding/json"
	"net/http"

	"github.com/raphael251/users-crud/pkg/utils"
)

func RespondBadRequest(w http.ResponseWriter, r *http.Request, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)

	errMessage := utils.Error{Message: "bad request", Data: data}
	json.NewEncoder(w).Encode(errMessage)
}

func RespondInternalServerError(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)

	errMessage := utils.Error{Message: "internal server error"}
	json.NewEncoder(w).Encode(errMessage)
}

func RespondNoContent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

func RespondNotFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
}

func RespondCreated(w http.ResponseWriter, r *http.Request, contentLocation string) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Location", contentLocation)
	w.WriteHeader(http.StatusCreated)
}

func RespondOK(w http.ResponseWriter, r *http.Request, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}
