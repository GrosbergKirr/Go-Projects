package main

import (
	"Wallet_intern/internal/api/v1/wallet"
	"Wallet_intern/internal/config"
	"Wallet_intern/internal/storage/postgressql"
	"github.com/go-chi/chi/v5"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	cfg := config.Mustload()

	log := setupLogger(cfg.Env)

	log.Info("Starting EWallet", slog.String("env", cfg.Env))
	log.Debug("debug messages are enabled")

	storage, err := postgressql.New(cfg.StoragePath)
	if err != nil {
		log.Error("failed to init storage. If you didn't set DB password, do it in local.yaml", err)
		os.Exit(1)
	}
	_ = storage

	router := chi.NewRouter()

	// --- Make create router ---
	router.Post("/api/v1/wallet", wallet.NewCreator(log, storage))

	// --- Make send router ---
	SendPattern := "/api/v1/wallet/{walletID}/send"
	router.Post(SendPattern, wallet.NewSender(log, storage))

	// --- Make transactions router ---
	HisPattern := "/api/v1/wallet/{walletID}/history"
	router.Get(HisPattern, wallet.NewHistoryGiver(log, storage))

	// --- Make statusget router ---
	StatusPattern := "/api/v1/wallet/{walletID}"
	router.Get(StatusPattern, wallet.NewStatusGetter(log, storage))

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
