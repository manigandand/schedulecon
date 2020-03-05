package main

import (
	"fmt"
	"hotstar/api"
	"hotstar/config"
	"hotstar/pkg/trace"
	"hotstar/store"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/rs/cors"
)

func main() {
	config.Initialize()
	trace.Setup(config.Env)
	store.Init()

	router := chi.NewRouter()
	cors := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "PATCH", "OPTIONS", "DELETE"},
		AllowedHeaders: []string{
			"Origin", "Authorization", "Access-Control-Allow-Origin",
			"Access-Control-Allow-Header", "Accept",
			"Content-Type", "X-CSRF-Token",
		},
		ExposedHeaders: []string{
			"Content-Length", "Access-Control-Allow-Origin", "Origin",
		},
		AllowCredentials: true,
		MaxAge:           300,
	})

	router.Use(cors.Handler)
	router.Use(
		middleware.Logger,
		middleware.Recoverer,
	)

	// Initialize the version 3 routes of the public API
	router.Route("/v1", api.Routes)

	trace.Log.Infof("Starting hotstar:1.0.0 on port :%s\n", config.Port)
	http.ListenAndServe(fmt.Sprintf(":%s", config.Port), router)
}
