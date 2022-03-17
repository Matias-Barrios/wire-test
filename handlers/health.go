package handlers

import (
	"fmt"
	"net/http"
)

type IHealthHandler interface {
	Health(w http.ResponseWriter, r *http.Request)
}

type HealthHandler struct {
}

func (h HealthHandler) Health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Ok!\n")
}

func NewHealthHandler() IHealthHandler {
	return HealthHandler{}
}
