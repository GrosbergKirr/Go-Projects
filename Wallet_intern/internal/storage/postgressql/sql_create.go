package postgressql

import (
	"log"
)

func (s *Storage) Create(idWall string, balanceWall float32) (int, error) {
	stmt, err := s.Db.Prepare("INSERT INTO wallet(id,balance) values($1,$2)")
	if err != nil {
		log.Fatalf("prepare mistake%s", err)
	}
	res, err := stmt.Exec(idWall, balanceWall)
	if err != nil {
		log.Fatalf("prepare mistake%s", err)
	}
	_ = res
	return 1, nil
}
