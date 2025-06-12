package routes

import (
	"fmt"
	"net/http"

	"github.com/devinmiller/fem-vanilla-js-go/internal/app"
	"github.com/go-chi/chi/v5"
)

func SetupRoutes(app *app.Application) *chi.Mux {
	r := chi.NewRouter()

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "status is available\n")
	})

	return r
}
