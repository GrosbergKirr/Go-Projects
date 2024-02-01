package wallet

import (
	"Wallet_intern/internal/tools"
	"errors"
	"github.com/go-chi/render"
	"io"
	"log/slog"
	"net/http"
)

type RequestSend struct {
	RecipientId string  `json:"id"`
	Amount      float32 `json:"amount"`
}

type Sender interface {
	Send(donorId string, recipientId string, amount float32) (int, error)
	CheckRecValid(recID string) (string, error)
	CheckDonorAmount(DonorId string) (float32, error)
}

func NewSender(log *slog.Logger, sender Sender, transDonorId string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		TransDonorId := transDonorId
		var req RequestSend

		err := render.DecodeJSON(r.Body, &req)

		if errors.Is(err, io.EOF) {
			// обработка ошибки с пустым запросом
			log.Error("request body is empty", http.StatusBadRequest)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if err != nil {
			log.Error("Transaction is failed!", http.StatusBadRequest)
		}

		RecipeValidCheck, err := sender.CheckRecValid(req.RecipientId)
		if err != nil {
			log.Error("RecipeValidCheck mistake")
		}
		DonorAmountCheck, err := sender.CheckDonorAmount(transDonorId)
		//Добавить проверку на тип!!!
		//(tools.TypeofObject(req.Amount) != "float32" не работает
		if (req.Amount < 0) || (tools.TypeofObject(req.Amount) != "float32") ||
			(RecipeValidCheck == "") ||
			(DonorAmountCheck < req.Amount) {

			log.Error("Transaction is failed!", http.StatusBadRequest)
			w.WriteHeader(http.StatusBadRequest)
			return
		} else {
			res, err := sender.Send(
				TransDonorId,
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
