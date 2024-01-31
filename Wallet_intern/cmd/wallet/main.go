package main

import (
	"Wallet_intern/internal/config"
	"Wallet_intern/internal/storage/psql"
	"fmt"
	"log/slog"
	"os"
)

func main() {
	cfg := config.Mustload()
	fmt.Println(cfg)

	log := setupLogger(cfg.Env)

	log.Info("Starting url-shortener", slog.String("env", cfg.Env))
	log.Debug("debug messages are enabled")

	storage, err := psql.New(cfg.StoragePath)
	if err != nil {
		log.Error("failed to init storage")
		os.Exit(1)
	}
	_ = storage

	id, err := storage.Send("1", "2", 100)
	if err != nil {
		log.Error("failed!!!")
	}
	_ = id
}

// ---Make logger---

// ---logger constant (several cases)---
const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}
	return log
}
