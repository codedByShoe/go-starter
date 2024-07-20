package app

import (
	"ABCD/internal/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func InitializeRouter() *chi.Mux {
	r := chi.NewRouter()
	// global middleware on all requests
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// mount sub routers from handlers
	r.Mount("/", handlers.HelloRoutes())

	return r
}
