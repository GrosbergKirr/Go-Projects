package psql

import "fmt"

// --- Check donor have enough money ---

func (s *Storage) CheckDonorAmount(DonorId string) (float32, error) {
	const op = "storage.psql.Check"

	type Rows struct {
		Amount float32
	}

	stmtCheck, err := s.Db.Prepare("select balance from wallet where id = $1")
	if err != nil {
		return 1, fmt.Errorf("%s: prepare statement: %w", op, err)
	}

	resCheck, err := stmtCheck.Query(DonorId)
	if err != nil {
		return 1, fmt.Errorf("%s: execute statement: %w", op, err)
	}
	p := Rows{}
	for resCheck.Next() {
		err := resCheck.Scan(&p.Amount)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
	return p.Amount, nil
}
