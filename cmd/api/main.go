package main

import (
	"card-validator/internal/env"
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

const version = "0.0.1"

func main() {
	cfg := config{
		addr:         env.GetString("PORT", ":8080"),
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	app := application{
		config: cfg,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))
}
