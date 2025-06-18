package app

import (
	"github.com/devinmiller/fem-vanilla-js-go/internal/handlers"
	"github.com/devinmiller/fem-vanilla-js-go/internal/logger"
)

type Application struct {
	Logger       *logger.Logger
	MovieHandler *handlers.MovieHandler
}

func NewApplication() (*Application, error) {
	// db connection will go here

	// logger will go here
	appLogger := logger.NewLogger()

	// stores will go here

	// handlers will go here
	movieHandler := handlers.NewMovieHandler(appLogger)

	app := &Application{
		Logger:       appLogger,
		MovieHandler: movieHandler,
	}

	return app, nil
}
