package psql

import (
	"fmt"
	"log"
	"time"
)

func (s *Storage) Send(donorId string, recipientId string, amount float32) (int, error) {
	const op = "storage.psql.Send"

	//--- SQL level money moving ---
	//--- Money + for recipient / Money - for donor ---
	stmtDonor, err := s.Db.Prepare("UPDATE wallet SET balance = balance - $1 WHERE id = $2")
	if err != nil {
		return 1, fmt.Errorf("%s: prepare statement: %w", op, err)
	}

	stmtRecipient, err := s.Db.Prepare("UPDATE wallet SET balance = balance + $1 WHERE id = $2")
	if err != nil {
		return 1, fmt.Errorf("%s: prepare statement: %w", op, err)
	}

	resDonor, err := stmtDonor.Exec(amount, donorId)
	if err != nil {
		return 1, fmt.Errorf("%s: execute statement: %w", op, err)
	}
	_ = resDonor

	resRecipient, err := stmtRecipient.Exec(amount, recipientId)
	if err != nil {
		return 1, fmt.Errorf("%s: execute statement: %w", op, err)
	}
	_ = resRecipient

	// --- Save transaction ---

	stmtTrans, err := s.Db.Prepare("INSERT INTO transactions (time, donor_id, recipient_id, amount ) values($1, $2, $3, $4)")
	if err != nil {
		return 1, fmt.Errorf("%s: prepare statement: %w", op, err)
	}
	resTrans, err := stmtTrans.Exec(time.Now(), donorId, recipientId, amount)
	if err != nil {
		log.Fatalf("problem: %s", err)
	}
	_ = resTrans
	return 1, nil

}
