package proxy

import (
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ProxyHandlerDeps struct {
	router *chi.Mux
	Pool   *pgxpool.Pool
}

type proxyHandler struct {
	deps *ProxyHandlerDeps
}

func New(deps *ProxyHandlerDeps) {
	proxyHandler := proxyHandler{
		deps: deps,
	}

	proxyHandler.deps.router.Route("/proxy", func(r chi.Router) {

	})
}
