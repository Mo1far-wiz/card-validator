package main

import (
	"card-validator/internal/api"
	"card-validator/internal/config"
	"log"
	"time"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println(err)
	}
}

const Version = "0.0.1"

func main() {
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
