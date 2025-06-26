package app

import (
	"database/sql"

	"github.com/devinmiller/fem-vanilla-js-go/internal/handlers"
	"github.com/devinmiller/fem-vanilla-js-go/internal/logger"
	"github.com/devinmiller/fem-vanilla-js-go/internal/store"
	"github.com/devinmiller/fem-vanilla-js-go/migrations"
)

type Application struct {
	Logger       *logger.Logger
	MovieHandler *handlers.MovieHandler
	DB           *sql.DB
}

func NewApplication() (*Application, error) {
	// db connection will go here
	pgDB, err := store.Open()
	if err != nil {
		return nil, err
	}

	err = store.MigrateFS(pgDB, migrations.FS, ".")
	if err != nil {
		panic(err)
	}

	// logger will go here
	appLogger := logger.NewLogger()

	// stores will go here
	movieStore := store.NewPostgresMovieStore(pgDB)

	// handlers will go here
	movieHandler := handlers.NewMovieHandler(movieStore, appLogger)

	app := &Application{
		Logger:       appLogger,
		MovieHandler: movieHandler,
		DB:           pgDB,
	}

	return app, nil
}
