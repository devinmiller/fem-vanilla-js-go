package app

import (
	"github.com/devinmiller/fem-vanilla-js-go/internal/logger"
)

type Application struct {
	Logger *logger.Logger
}

func NewApplication() (*Application, error) {
	// db connection will go here

	// logger will go here
	appLogger := logger.NewLogger()

	// stores will go here

	// handlers will go here

	app := &Application{
		Logger: appLogger,
	}

	return app, nil
}
