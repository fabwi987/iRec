package models

import _ "github.com/go-sql-driver/mysql"

type Position struct {
	Id     int
	Userid *User
	Title  string
	Body   string
	Reward string
}

func (db *DB) GetPositions() ([]*Position, error) {

	rows, err := db.Query("SELECT * FROM positions")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	poss := make([]*Position, 0)
	var id int
	for rows.Next() {
		ps := new(Position)
		err = rows.Scan(&ps.Id, id, &ps.Title, &ps.Body, &ps.Reward)
		if err != nil {
			return nil, err
		}

		ps.Userid, err = db.GetUser(id)
		if err != nil {
			return nil, err
		}

		poss = append(poss, ps)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return poss, nil
}

func (db *DB) GetPosition(positionid int) (*Position, error) {

	stmt, err := db.Prepare("SELECT * FROM positions WHERE idpositions = ?")
	defer stmt.Close()
	rows, err := stmt.Query(positionid)
	defer rows.Close()
	ps := new(Position)
	var id int

	for rows.Next() {
		err = rows.Scan(&ps.Id, id, &ps.Title, &ps.Body, &ps.Reward)
		if err != nil {
			return nil, err
		}

		ps.Userid, err = db.GetUser(id)
		if err != nil {
			return nil, err
		}

	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return ps, nil

}
