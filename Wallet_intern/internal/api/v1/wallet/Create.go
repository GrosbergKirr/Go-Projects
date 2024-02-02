package wallet

import (
	"Wallet_intern/internal/tools"
	"github.com/go-chi/render"
	"log/slog"
	"net/http"
)

type ResponseCreate struct {
	Id      string  `json:"id"`
	Balance float32 `json:"balance"`
}

type Creator interface {
	Create(idWall string, balanceWall float32) (int, error)
}

func NewCreator(log *slog.Logger, creator Creator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		newId := tools.NewRandomString(30)
		var defaultMoney float32 = 100

		id, err := creator.Create(
			newId,
			defaultMoney)

		if err != nil {
			log.Error("failed to create wallet!")
		}

		log.Info("Wallet created. ID:", id)

		CreateRespOK(w, r, newId, defaultMoney)
	}
}
func CreateRespOK(w http.ResponseWriter, r *http.Request, id string, balance float32) {
	render.JSON(w, r, ResponseCreate{
		Id:      id,
		Balance: balance,
	})
}
