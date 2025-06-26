package handlers

import (
	"net/http"

	"github.com/devinmiller/fem-vanilla-js-go/internal/logger"
	"github.com/devinmiller/fem-vanilla-js-go/internal/store"
	"github.com/devinmiller/fem-vanilla-js-go/internal/utils"
)

type MovieHandler struct {
	movieStore store.MovieStore
	logger     *logger.Logger
}

func NewMovieHandler(movieStore store.MovieStore, logger *logger.Logger) *MovieHandler {
	return &MovieHandler{
		movieStore: movieStore,
		logger:     logger,
	}
}

func (m *MovieHandler) GetTopMovies(w http.ResponseWriter, r *http.Request) {
	movies, err := m.movieStore.GetTopMovies()
	if err != nil {
		m.logger.Error("GetTopMovies - Query", err)
		utils.WriteError(w, http.StatusInternalServerError)
		return
	}

	err = utils.WriteJSON(w, http.StatusOK, utils.Envelope{"movies": movies})
	if err != nil {
		m.logger.Error("GetTopMovies - Encoding", err)
		utils.WriteError(w, http.StatusInternalServerError)
	}
}

func (m *MovieHandler) GetRandomMovies(w http.ResponseWriter, r *http.Request) {
	movies, err := m.movieStore.GetRandomMovies()
	if err != nil {
		m.logger.Error("GetRandomMovies - Query", err)
		utils.WriteError(w, http.StatusInternalServerError)
		return
	}

	err = utils.WriteJSON(w, http.StatusOK, utils.Envelope{"movies": movies})
	if err != nil {
		m.logger.Error("GetRandomMovies - Encoding", err)
		utils.WriteError(w, http.StatusInternalServerError)
	}
}
