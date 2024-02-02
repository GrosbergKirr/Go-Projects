package postgressql

import "fmt"

// --- Check recipientId is valid ---

func (s *Storage) CheckRecValid(recID string) (string, error) {
	const op = "storage.postgressql.Check"

	type Rows struct {
		Id string
	}

	stmtCheck, err := s.Db.Prepare("select id from wallet where id = $1")
	if err != nil {
		return "1", fmt.Errorf("%s: prepare statement: %w", op, err)
	}

	resCheck, err := stmtCheck.Query(recID)
	if err != nil {
		return "1", fmt.Errorf("%s: execute statement: %w", op, err)
	}
	p := Rows{}
	for resCheck.Next() {
		err := resCheck.Scan(&p.Id)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
	return p.Id, nil
}
