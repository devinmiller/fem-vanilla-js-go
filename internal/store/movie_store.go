package store

import "github.com/devinmiller/fem-vanilla-js-go/internal/models"

type MovieStore interface {
	GetTopMovies() ([]models.Movie, error)
	GetRandomMovies() ([]models.Movie, error)
	GetMovieById(id int) (models.Movie, error)
	SearchMoviesByName(name string) ([]models.Movie, error)
	GetAllGenres() ([]models.Genre, error)
}
