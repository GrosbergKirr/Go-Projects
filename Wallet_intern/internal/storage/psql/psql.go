package psql

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"time"
)

type Storage struct {
	db *sql.DB
}

func New(storagePath string) (*Storage, error) {
	const op = "storage.psql.New"

	db, err := sql.Open("postgres", storagePath)

	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	cr, err := db.Exec("create table IF NOT EXISTS wallet (id varchar, balance float4, PRIMARY KEY (id))")
	cr1, err := db.Exec("create table IF NOT EXISTS transactions (time timestamp ,donor_id varchar, recipient_id varchar, amount float4)")

	_ = cr
	_ = cr1
	if err != nil {
		return nil, fmt.Errorf("{op}: #{err}")
	}
	return &Storage{db: db}, nil
}

func (s *Storage) Create(idWall string, balanceWall float32) (int, error) {
	const op = "storage.psql.Create"

	stmt, err := s.db.Prepare("INSERT INTO wallet(id,balance) values($1,$2)")
	if err != nil {
		return 0, fmt.Errorf("%s: prepare statement: %w", op, err)
	}
	res, err := stmt.Exec(idWall, balanceWall)
	if err != nil {
		return 0, fmt.Errorf("problem: %s", err)
	}
	_ = res
	return 0, nil
}

// --- Method SEND ---

func (s *Storage) Send(donorId string, recipientId string, amount float32) (int, error) {
	const op = "storage.psql.Send"

	//--- SQL level money moving ---
	//--- Money + for recipient / Money - for donor ---
	stmtDonor, err := s.db.Prepare("UPDATE wallet SET balance = balance - $1 WHERE id = $2")
	if err != nil {
		return 0, fmt.Errorf("%s: prepare statement: %w", op, err)
	}

	stmtRecipient, err := s.db.Prepare("UPDATE wallet SET balance = balance + $1 WHERE id = $2")
	if err != nil {
		return 0, fmt.Errorf("%s: prepare statement: %w", op, err)
	}

	resDonor, err := stmtDonor.Exec(amount, donorId)
	if err != nil {
		return 0, fmt.Errorf("%s: execute statement: %w", op, err)
	}
	_ = resDonor

	resRecipient, err := stmtRecipient.Exec(amount, recipientId)
	if err != nil {
		return 0, fmt.Errorf("%s: execute statement: %w", op, err)
	}
	_ = resRecipient

	// --- Save transaction ---

	stmtTrans, err := s.db.Prepare("INSERT INTO transactions (time, donor_id, recipient_id, amount ) values($1, $2, $3, $4)")
	if err != nil {
		return 0, fmt.Errorf("%s: prepare statement: %w", op, err)
	}
	resTrans, err := stmtTrans.Exec(time.Now(), donorId, recipientId, amount)
	if err != nil {
		log.Fatalf("problem: %s", err)
	}
	_ = resTrans
	return 0, nil
}
