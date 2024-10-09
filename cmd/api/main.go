package main

import (
	"log"
	"social/internal/env"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	cfg := config{
		address: env.GetString("ADDRESS", ":3000"),
	}

	app := application{
		config: cfg,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))
}
