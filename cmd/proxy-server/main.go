package main

import (
	"fmt"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"proxyStoreServer/internal/app"
	"proxyStoreServer/internal/config"
	"syscall"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	config := config.New()
	logger := setupLogger(&config.Application)

	sigChannel := make(chan os.Signal, 1)
	signal.Notify(sigChannel, syscall.SIGINT, syscall.SIGTERM)

	app, closeFn := app.New(config, logger)

	go func() {
		<-sigChannel
		fmt.Println("Shutting down server...")
		closeFn()
	}()

	log.Fatal(app.Run())

}

func setupLogger(applicationConfig *config.ApplicationConfig) *slog.Logger {

	var logger *slog.Logger

	switch applicationConfig.Environment {
	case "local", "dev":
		logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		}))
	case "production":
		logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelInfo,
		}))
	}

	return logger
}
