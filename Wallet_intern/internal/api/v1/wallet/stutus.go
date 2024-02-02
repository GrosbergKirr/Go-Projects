package wallet

import (
	"Wallet_intern/internal/storage/postgressql"
	"github.com/go-chi/render"
	"log/slog"
	"net/http"
)

type ResponseStatus struct {
	Id      string  `json:"id"`
	Balance float32 `json:"balance"`
}

type StatusGetter interface {
	WalletGetter(WalletId string) (postgressql.Wallet, error)
}

func NewStatusGetter(log *slog.Logger, wallgetter StatusGetter, walletId string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		wallet, err := wallgetter.WalletGetter(walletId)
		if err != nil {
			log.Error("Get wallet mistake", err)
		}
		if wallet.Id == "" {
			log.Error("The wallet does not exist!")
			w.WriteHeader(http.StatusNotFound)
		}
		StatusGetRespOK(w, r, wallet.Id, wallet.Amount)
	}
}

func StatusGetRespOK(w http.ResponseWriter, r *http.Request, id string, balance float32) {
	render.JSON(w, r, ResponseStatus{
		Id:      id,
		Balance: balance,
	})
}
