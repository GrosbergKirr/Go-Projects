package postgressql

import (
	"fmt"
	"log"
	"time"
)

type WalletHistory struct {
	Time   time.Time
	From   string
	To     string
	Amount float32
}

func (s *Storage) History(walletId string) ([]WalletHistory, error) {

	hisStmt, err := s.Db.Prepare("select * from transactions where (donor_id = $1 or recipient_id = $1)")
	if err != nil {
		log.Fatalf("prepare mistake%s", err)
	}

	history, err := hisStmt.Query(walletId)
	if err != nil {
		log.Fatalf("query mistake%s", err)
	}
	var transactions []WalletHistory

	for history.Next() {
		wh := WalletHistory{}
		err := history.Scan(&wh.Time, &wh.From, &wh.To, &wh.Amount)
		if err != nil {
			fmt.Println(err)
			continue
		}
		transactions = append(transactions, wh)
	}
	return transactions, nil
}
