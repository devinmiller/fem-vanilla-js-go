package utils

import (
	"encoding/json"
	"net/http"
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
