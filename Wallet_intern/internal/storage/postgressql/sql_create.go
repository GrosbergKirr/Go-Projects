package postgressql

import "fmt"

func (s *Storage) Create(idWall string, balanceWall float32) (int, error) {
	const op = "storage.postgressql.Create"

	stmt, err := s.Db.Prepare("INSERT INTO wallet(id,balance) values($1,$2)")
	if err != nil {
		return 1, fmt.Errorf("%s: prepare statement: %w", op, err)
	}
	res, err := stmt.Exec(idWall, balanceWall)
	if err != nil {
		return 1, fmt.Errorf("problem: %s", err)
	}
	_ = res
	return 1, nil
}
