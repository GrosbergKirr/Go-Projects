package wallet

import (
	"Wallet_intern/internal/storage/postgressql"
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"io"
	"log/slog"
	"net/http"
)

type RequestSend struct {
	RecipientId string  `json:"to"`
	Amount      float32 `json:"amount"`
}

type Sender interface {
	Send(donorId string, recipientId string, amount float32) (int, error)
	WalletGetter(WalletId string) (postgressql.Wallet, error)
}

func NewSender(log *slog.Logger, sender Sender) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		WalletId := chi.URLParam(r, "walletID")

		var req RequestSend

		err := render.DecodeJSON(r.Body, &req)

		if errors.Is(err, io.EOF) {
			// обработка ошибки с пустым запросом
			log.Error("request body is empty", http.StatusBadRequest)
			w.WriteHeader(http.StatusBadRequest)
		}
		if err != nil {
			log.Error("Transaction is failed!", http.StatusBadRequest)
			w.WriteHeader(http.StatusBadRequest)
		} else {
			RecipeValidCheck, err := sender.WalletGetter(req.RecipientId)
			DonorAmountCheck, err := sender.WalletGetter(WalletId)
			if err != nil {
				log.Error("get id from DB mistake", http.StatusBadRequest)
			}
			if req.Amount < 0 {
				log.Error("Requesting money < 0", http.StatusBadRequest)
				w.WriteHeader(http.StatusBadRequest)
			} else if req.RecipientId == "" {
				log.Error("Empty recipe id", http.StatusBadRequest)
				w.WriteHeader(http.StatusBadRequest)

			} else if DonorAmountCheck.Amount < req.Amount {
				log.Error("Donor haven't enough money", http.StatusBadRequest)
				w.WriteHeader(http.StatusBadRequest)
			} else if RecipeValidCheck.Id == "" {
				log.Error("Couldn't find recipient wallet", http.StatusBadRequest)
				w.WriteHeader(http.StatusBadRequest)
			} else {
				res, err := sender.Send(
					WalletId,
					req.RecipientId,
					req.Amount,
				)
				if err != nil {
					log.Error("failed to create wallet!")
				}
				_ = res
				log.Info("Transaction success!")

			}
		}

	}
}
