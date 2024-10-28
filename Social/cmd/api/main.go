package main

import (
	"Social/cmd/api/internal/env"
	"log"
	"os"
)

func main() {
	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
	}

	app := &application{
		config: cfg,
	}

	os.LookupEnv("PATH")

	mux := app.mount()
	log.Fatal(app.run(mux))
}
