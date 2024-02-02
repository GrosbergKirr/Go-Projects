package wallet

import (
	"Wallet_intern/internal/storage/postgressql"
	"github.com/go-chi/render"
	"log/slog"
	"net/http"
)

type ResponseHistory struct {
	WalletTrans []postgressql.WalletHistory `json:"transactions"`
}

type HistoryGiver interface {
	History(walletId string) ([]postgressql.WalletHistory, error)
}

func NewHistoryGiver(log *slog.Logger, HisGiv HistoryGiver, WalletId string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		trans, err := HisGiv.History(WalletId)
		if err != nil {
			log.Error("Ошибка получение данных из таблицы переводов:", err)
			w.WriteHeader(http.StatusBadRequest)
		}
		log.Info("wallet transaction history is available")
		HistoryRespOK(w, r, trans)

	}
}

func HistoryRespOK(w http.ResponseWriter, r *http.Request, trans []postgressql.WalletHistory) {
	render.JSON(w, r, ResponseHistory{
		WalletTrans: trans,
	})
}
