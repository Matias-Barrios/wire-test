//go:build wireinject
// +build wireinject

package main

import (
	"github.com/Matias-Barrios/wire-test/config"
	"github.com/Matias-Barrios/wire-test/handlers"
	"github.com/google/wire"
)

func InitializeApp() (App, error) {
	wire.Build(
		NewApp,
		config.LoadConfig,
		handlers.NewHealthHandler,
		handlers.NewRouter)
	return App{}, nil
}
