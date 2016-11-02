package models

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

//Position is the representation of a job position
type Position struct {
	Id     int
	Userid *User
	Title  string
	Body   string
	Reward string
}

//GetPositions return all positions from the database
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
		err = rows.Scan(&ps.Id, &id, &ps.Title, &ps.Body, &ps.Reward)
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

//GetPosition returns a single position from it's id
func (db *DB) GetPosition(positionid int) (*Position, error) {

	stmt, err := db.Prepare("SELECT * FROM positions WHERE idpositions = ?")
	defer stmt.Close()
	rows, err := stmt.Query(positionid)
	defer rows.Close()
	ps := new(Position)
	var id int

	for rows.Next() {
		err = rows.Scan(&ps.Id, &id, &ps.Title, &ps.Body, &ps.Reward)
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

//CreatePosition creates a new position inte the database
func (db *DB) CreatePosition(NewPosition *Position) error {

	stmt, err := db.Prepare("INSERT positions SET iduser=?,title=?,body=?,reward=?")
	if err != nil {
		return err
	}

	res, err := stmt.Exec(NewPosition.Userid.Id, NewPosition.Title, NewPosition.Body, NewPosition.Reward)
	if err != nil {
		return err
	}
	fmt.Println(res)
	return nil
}
