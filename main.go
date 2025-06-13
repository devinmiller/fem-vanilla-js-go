package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/devinmiller/fem-vanilla-js-go/internal/app"
	"github.com/devinmiller/fem-vanilla-js-go/internal/routes"
)

func main() {
	var port int
	flag.IntVar(&port, "port", 8081, "go backend server port")
	flag.Parse()

	app, err := app.NewApplication()
	if err != nil {
		panic(err)
	}

	r := routes.SetupRoutes(app)
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      r,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	err = server.ListenAndServe()
	if err != nil {
		app.Logger.Fatal("error starting server", err)
	}
}
