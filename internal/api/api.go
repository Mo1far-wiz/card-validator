package api

import (
	"card-validator/internal/api/handlers"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type Application struct {
	Config Config
}

type Config struct {
	Addr         string
	WriteTimeout time.Duration
	ReadTimeout  time.Duration
	IdleTimeout  time.Duration
}

func (app *Application) Run(mux http.Handler) error {
	srv := http.Server{
		Addr:         app.Config.Addr,
		Handler:      mux,
		WriteTimeout: app.Config.WriteTimeout,
		ReadTimeout:  app.Config.ReadTimeout,
		IdleTimeout:  app.Config.IdleTimeout,
	}

	log.Printf("server has started at %s", app.Config.Addr)

	return srv.ListenAndServe()
}

func (app *Application) Mount() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Timeout(time.Minute))

	r.Route("/v1", func(r chi.Router) {
		r.Get("/health", handlers.HealthCheckHandler)
		r.Post("/validate-card", handlers.ValidateCardHandler)
	})

	return r
}
