package main

import (
	"log"

	"github.com/otaviouu/go_social/internal/env"
	"github.com/otaviouu/go_social/internal/store"
)

func main() {
	cfg := &config{
		addr: env.GetString("ADDR", ":8080"),
	}
	storage := store.NewStorage(nil)

	app := &application{
		config:  *cfg,
		storage: storage,
	}

	router := app.mount()
	log.Fatal(app.run(router))
}
