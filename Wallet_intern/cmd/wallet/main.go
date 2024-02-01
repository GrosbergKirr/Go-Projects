package main

import (
	"Wallet_intern/internal/api/v1/wallet"
	"Wallet_intern/internal/config"
	"Wallet_intern/internal/storage/psql"
	"github.com/go-chi/chi/v5"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	cfg := config.Mustload()

	log := setupLogger(cfg.Env)

	log.Info("Starting url-shortener", slog.String("env", cfg.Env))
	log.Debug("debug messages are enabled")

	storage, err := psql.New(cfg.StoragePath)
	if err != nil {
		log.Error("failed to init storage")
		os.Exit(1)
	}
	_ = storage

	router := chi.NewRouter()

	router.Post("/api/v1/wallet", wallet.NewCreator(log, storage))

	donorid := "tzc7ny5i96qssbdb97xuv02deacw5k"
	pattern := "/api/v1/wallet/" + donorid + "/send"
	router.Post(pattern, wallet.NewSender(log, storage, donorid))

	//ЗАПУСКАЕМ СЕРВЕР !!!!

	log.Info("starting server", slog.String("address", cfg.Address))

	srv := &http.Server{
		Addr:         cfg.Address,
		Handler:      router,
		ReadTimeout:  cfg.HTTPServer.Timeout,
		WriteTimeout: cfg.HTTPServer.Timeout,
		IdleTimeout:  cfg.HTTPServer.IdleTimeout,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Error("failed to start server")
	}

	log.Info("server started")
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