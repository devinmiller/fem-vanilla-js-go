package routes

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/devinmiller/fem-vanilla-js-go/internal/app"
	"github.com/go-chi/chi/v5"
)

func SetupRoutes(app *app.Application) *chi.Mux {
	r := chi.NewRouter()

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "status is available\n")
	})

	r.Get("/api/movies/top", app.MovieHandler.GetTopMovies)
	r.Get("/api/movies/random", app.MovieHandler.GetRandomMovies)
	r.Get("/api/movies/{id}", app.MovieHandler.GetMovie)

	clientRouteHandler := func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filepath.Join("public", "index.html"))
	}

	// TODO: Find a more dynamic method for handling client routing
	r.Get("/movies/{id}", clientRouteHandler)
	r.Get("/account", clientRouteHandler)
	r.Get("/account/favorites", clientRouteHandler)
	r.Get("/account/watchlist", clientRouteHandler)

	filesDir := http.Dir("public")
	FileServer(r, "/", filesDir)

	return r
}

func FileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit any URL parameters")
	}

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", http.StatusPermanentRedirect).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		rctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
		fs := http.StripPrefix(pathPrefix, http.FileServer(root))
		fs.ServeHTTP(w, r)
	})
}
