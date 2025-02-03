package main

import (
	"card-validator/internal/api"
	"card-validator/internal/api/handlers"
	"card-validator/internal/config"
	"card-validator/internal/domain/validator"
	json "card-validator/internal/utils/json"
	"log"
	"os"
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

	log.Fatal(app.Run(mux))
}
