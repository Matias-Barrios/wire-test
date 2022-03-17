package handlers

import (
	"github.com/gorilla/mux"
)

func NewRouter(healthHandler IHealthHandler) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", healthHandler.Health)
	return r
}
