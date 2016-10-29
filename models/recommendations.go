package models

import _ "github.com/go-sql-driver/mysql"

type Recommendation struct {
	ID         int
	IdPosition *Position
	IdUser     *User
	IdUser2    *User
}

func (db *DB) GetRecommendations() ([]*Recommendation, error) {

	rows, err := db.Query("SELECT * FROM stocks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	recs := make([]*Recommendation, 0)
	var idPos int
	var idUsr int
	var idUsr2 int

	for rows.Next() {
		rec := new(Recommendation)
		err = rows.Scan(rec.ID, idPos, idUsr, idUsr2)
		if err != nil {
			return nil, err
		}

		rec.IdPosition, err = db.GetPosition(idPos)
		if err != nil {
			return nil, err
		}

		rec.IdUser, err = db.GetUser(idUsr)
		if err != nil {
			return nil, err
		}

		rec.IdUser2, err = db.GetUser(idUsr2)
		if err != nil {
			return nil, err
		}

	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return recs, nil
}

func (db *DB) GetRecommendation(recommendationid int) (*Recommendation, error) {

	stmt, err := db.Prepare("SELECT * FROM recommendations WHERE idrecommendations = ?")
	defer stmt.Close()
	rows, err := stmt.Query(recommendationid)
	defer rows.Close()
	rec := new(Recommendation)
	var idPos int
	var idUsr int
	var idUsr2 int

	for rows.Next() {

		err = rows.Scan(rec.ID, idPos, idUsr, idUsr2)
		if err != nil {
			return nil, err
		}

		rec.IdPosition, err = db.GetPosition(idPos)
		if err != nil {
			return nil, err
		}

		rec.IdUser, err = db.GetUser(idUsr)
		if err != nil {
			return nil, err
		}

		rec.IdUser2, err = db.GetUser(idUsr2)
		if err != nil {
			return nil, err
		}

	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return rec, nil
}
