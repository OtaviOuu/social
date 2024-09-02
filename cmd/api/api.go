package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/otaviouu/go_social/internal/store"
)

type application struct {
	config  config
	storage store.Storage
}

type config struct {
	addr string
}

// Monta o Router
func (app *application) mount() http.Handler {

	// Router ~ Multiplexer
	r := chi.NewRouter()

	// Middlewares
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))

	r.Route("/v1", func(r chi.Router) {
		r.Get("/health", app.healthCheckerHandler)
	})

	return r
}

func (app *application) run(router http.Handler) error {

	// Instanciando um Server antes
	// para facilitar a configuracao
	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      router,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		//IdleTimeout:  time.Minute,
	}

	log.Print("Server is running at port: ", srv.Addr[1:])
	return srv.ListenAndServe()
}
