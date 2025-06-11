package app

import (
	"log"
	"os"
)

type Application struct {
	Logger *log.Logger
}

func NewApplication() (*Application, error) {
	// db connection will go here

	// logger will go here
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	// stores will go here

	// handlers will go here

	app := &Application{
		Logger: logger,
	}

	return app, nil
}
