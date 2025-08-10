package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/lib/pq"
	"github.com/thetnaingtn/tidy-url/internal/config"
	"github.com/thetnaingtn/tidy-url/server"
	"github.com/thetnaingtn/tidy-url/store"
	"github.com/thetnaingtn/tidy-url/store/db"
)

func main() {
	config := &config.Config{
		Addr:    getEnv("ADDR", ""),
		DSN:     getEnv("DB_CONNECTION_URL", "postgres://tidyurl:pa55w0rd@localhost:5432/tidyurl?sslmode=disable"),
		BaseURL: getEnv("BASE_URL", "http://localhost:8080"),
		Port:    getEnv("PORT", "8080"),
	}

	driver, err := db.NewDBDriver(config)

	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithCancel(context.Background())

	store := store.NewStore(config, driver)

	s, err := server.NewServer(ctx, store, config)
	if err != nil {
		cancel()
		slog.Error("failed to create server", "error", err)
	}

	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	if err := s.Start(); err != nil {
		if err != http.ErrServerClosed {
			cancel()
			slog.Error("failed to start server", "error", err)
		}
	}

	go func() {
		<-c
		cancel()
	}()

	<-ctx.Done()
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
