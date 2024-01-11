package handlers

import (
	"net/http"
)

type ApplicationHandler struct{}

func NewApplicationHandler() *ApplicationHandler {
	return &ApplicationHandler{}
}

func (h *ApplicationHandler) Health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}
