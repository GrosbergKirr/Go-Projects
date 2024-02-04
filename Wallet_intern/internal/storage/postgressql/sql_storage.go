package postgressql

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type Storage struct {
	Db *sql.DB
}

func New(storagePath string) (*Storage, error) {

	db, err := sql.Open("postgres", storagePath)

	if err != nil {
		return nil, err
	}

	cr, err := db.Exec("create table IF NOT EXISTS wallet (id varchar, balance float4, PRIMARY KEY (id))")
	cr1, err := db.Exec("create table IF NOT EXISTS transactions (time timestamp ,donor_id varchar, recipient_id varchar, amount float4)")

	_ = cr
	_ = cr1
	if err != nil {
		return nil, err
	}
	return &Storage{Db: db}, nil
}
