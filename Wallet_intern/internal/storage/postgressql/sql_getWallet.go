package postgressql

import (
	"fmt"
	"log"
)

type Wallet struct {
	Id     string
	Amount float32
}

func (s *Storage) WalletGetter(WalletId string) (Wallet, error) {
	const op = "storage.postgressql.Check"

	stmtCheck, err := s.Db.Prepare("select * from wallet where id = $1")
	if err != nil {
		log.Fatalf("prepare mistake%s", err)
	}

	resCheck, err := stmtCheck.Query(WalletId)
	if err != nil {
		log.Fatalf("prepare mistake%s", err)
	}
	wallet := Wallet{}
	for resCheck.Next() {
		err := resCheck.Scan(&wallet.Id, &wallet.Amount)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
	return wallet, nil
}
