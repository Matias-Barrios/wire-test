package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Matias-Barrios/wire-test/config"
	"github.com/gorilla/mux"
)

type App struct {
	Config *config.Config
	Router *mux.Router
}

func (a App) Start() {
	srv := &http.Server{
		Handler:      a.Router,
		Addr:         fmt.Sprintf("127.0.0.1:%d", a.Config.Port),
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  5 * time.Second,
	}
	log.Println(fmt.Sprintf("Starting server http on port %d", a.Config.Port))
	log.Fatal(srv.ListenAndServe())
}

func NewApp(c *config.Config, r *mux.Router) (App, error) {
	return App{
		Config: c,
		Router: r,
	}, nil
}
