package utils

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type Envelope map[string]any

func WriteJSON(w http.ResponseWriter, status int, data Envelope) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(jsonData)

	return nil
}

func WriteError(w http.ResponseWriter, status int) error {
	var msg string
	switch status {
	case http.StatusBadRequest:
		msg = "invalid request"
	case http.StatusInternalServerError:
		fallthrough
	default:
		msg = "internal server error"
	}

	data := &Envelope{"error": msg}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(jsonData)

	return nil
}

func ReadIDParam(r *http.Request) (int, error) {
	idParam := chi.URLParam(r, "id")
	if idParam == "" {
		return 0, errors.New("invalid id parameter")
	}

	id, err := strconv.Atoi(idParam)
	if err != nil {
		return 0, errors.New("invalid id parameter type")
	}

	return id, nil
}
