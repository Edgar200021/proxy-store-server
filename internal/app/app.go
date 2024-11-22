package app

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"proxyStoreServer/internal/bot_client"
	"proxyStoreServer/internal/config"
	"proxyStoreServer/internal/cryptomus_client"
	"proxyStoreServer/internal/router"
	"proxyStoreServer/internal/storage/postgres"
)

type App struct {
	port   uint32
	server *http.Server
}

func (app *App) Run() error {
	fmt.Println("Running on port", app.port)

	return app.server.ListenAndServe()

}

func (app *App) Port() uint32 {
	return app.port

}

func New(config *config.Config, logger *slog.Logger) (*App, func()) {
	var (
		botClient = bot_client.New(&config.Bot, &config.Application)
		_         = cryptomus_client.New(&config.Cryptomus)
	)

	if err := botClient.SetWebHook(); err != nil {
		log.Fatal("Failed to set webhook ", err)
	}

	db, err := postgres.New(&config.Database)
	if err != nil {
		log.Fatal("Failed to connect to database ", err)
	}

	mux := router.New(config)

	server := http.Server{
		Addr:         fmt.Sprintf("%s:%d", config.Application.Host, config.Application.Port),
		Handler:      mux,
		WriteTimeout: config.Application.WriteTimeout,
		ReadTimeout:  config.Application.ReadTimeout,
		IdleTimeout:  config.Application.IdleTimeout,
	}

	closeFn := func() {
		db.Close()
		server.Shutdown(context.Background())
	}

	return &App{
		port:   config.Application.Port,
		server: &server,
	}, closeFn
}
