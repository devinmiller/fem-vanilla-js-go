package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/devinmiller/fem-vanilla-js-go/internal/logger"
	"github.com/devinmiller/fem-vanilla-js-go/internal/models"
	"github.com/devinmiller/fem-vanilla-js-go/internal/utils"
)

type MovieHandler struct {
	logger *logger.Logger
}

func NewMovieHandler(logger *logger.Logger) *MovieHandler {
	return &MovieHandler{
		logger: logger,
	}
}

func (m *MovieHandler) GetTopMovies(w http.ResponseWriter, r *http.Request) {
	movies := []models.Movie{
		{
			ID:          1,
			TMDB_ID:     181,
			Title:       "The Hacker",
			ReleaseYear: 2022,
			Genres:      []models.Genre{{ID: 1, Name: "Thriller"}},
			Keywords:    []string{},
			Casting:     []models.Actor{{ID: 1, FirstName: "Max"}},
		},
		{
			ID:          2,
			TMDB_ID:     181,
			Title:       "Back to the Future",
			ReleaseYear: 1984,
			Genres:      []models.Genre{{ID: 1, Name: "Thriller"}},
			Keywords:    []string{},
			Casting:     []models.Actor{{ID: 1, FirstName: "Max"}},
		},
	}

	if err := json.NewEncoder(w).Encode(movies); err != nil {
		m.logger.Error("encoding GetTopMovies: %v", err)
		utils.WriteError(w, http.StatusInternalServerError)
		return
	}

	utils.WriteJSON(w, http.StatusOK, utils.Envelope{"movies": movies})
}

func (m *MovieHandler) GetRandomMovies(w http.ResponseWriter, r *http.Request) {
	movies := []models.Movie{
		{
			ID:          1,
			TMDB_ID:     181,
			Title:       "The Random Hacker",
			ReleaseYear: 2022,
			Genres:      []models.Genre{{ID: 1, Name: "Thriller"}},
			Keywords:    []string{},
			Casting:     []models.Actor{{ID: 1, FirstName: "Max"}},
		},
		{
			ID:          2,
			TMDB_ID:     181,
			Title:       "Back to the Random Future",
			ReleaseYear: 1984,
			Genres:      []models.Genre{{ID: 1, Name: "Thriller"}},
			Keywords:    []string{},
			Casting:     []models.Actor{{ID: 1, FirstName: "Max"}},
		},
	}

	if err := json.NewEncoder(w).Encode(movies); err != nil {
		m.logger.Error("encoding GetRandomMovies: %v", err)
		utils.WriteError(w, http.StatusInternalServerError)
		return
	}

	utils.WriteJSON(w, http.StatusOK, utils.Envelope{"movies": movies})
}
