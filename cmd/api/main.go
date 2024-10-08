package main

import (
	"log"
	"social/internal/env"
)

func main() {
	cfg := config{
		address: env.GetString("ADDRESS", ":3000"),
	}

	app := application{
		config: cfg,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))
}
