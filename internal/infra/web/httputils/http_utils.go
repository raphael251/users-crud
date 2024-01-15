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