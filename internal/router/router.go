package router

import (
	"net/http"
	"proxyStoreServer/internal/config"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func New(config *config.Config) *chi.Mux {
	router := chi.NewRouter()

	apiRouter := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.URLFormat)
	router.Use(middleware.RequestID)
	router.Use(middleware.Recoverer)
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{config.Application.ClientUrl},
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowCredentials: true,
	}))

	router.Mount("/api", apiRouter)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Test"))
	})

	return router
}
