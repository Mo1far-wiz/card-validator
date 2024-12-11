package main

import (
	"card-validator/cmd/controllers"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type application struct {
	config config
}

type config struct {
	addr         string
	WriteTimeout time.Duration
	ReadTimeout  time.Duration
	IdleTimeout  time.Duration
}

func (app *application) run(mux http.Handler) error {
	srv := http.Server{
		Addr:         app.config.addr,
		Handler:      mux,
		WriteTimeout: app.config.WriteTimeout,
		ReadTimeout:  app.config.ReadTimeout,
		IdleTimeout:  app.config.IdleTimeout,
	}

	log.Printf("server has started at %s", app.config.addr)

	return srv.ListenAndServe()
}

func (app *application) mount() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Timeout(time.Minute))

	r.Route("/v1", func(r chi.Router) {
		r.Get("/health", controllers.HealthCheckHandler)
		r.Post("/validate-card", controllers.ValidateCardHandler)
	})

	return r
}
