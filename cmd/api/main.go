package main

import (
	"card-validator/internal/api"
	"card-validator/internal/api/handlers"
	"card-validator/internal/config"
	"card-validator/internal/domain/validator"
	json "card-validator/internal/utils/json"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	json_validator "github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

const Version = "0.0.1"

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println(err)
	}

	logger := log.New(os.Stdout, "Custom INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	v := &validator.CreditCardValidator{
		Logger: logger,
	}
	handlers.CardValidator = validator.NewCardValidator(v)

	json.Validate = json_validator.New(json_validator.WithRequiredStructEnabled())

	cfg := api.Config{
		Addr:         config.GetString("PORT", ":8080"),
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	app := api.Application{
		Config: cfg,
	}

	mux := app.Mount()

	// Graceful Shutdown
	server := http.Server{
		Addr:    app.Config.Addr,
		Handler: mux,
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	log.Print("Server Started")

	<-done
	log.Print("Server Stopped")

	// shutdown context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Gracefully shutdown
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown Failed:%+v", err)
	}
	log.Print("Server Exited Properly")
}
