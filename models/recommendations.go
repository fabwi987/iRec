package models

import _ "github.com/go-sql-driver/mysql"

//Recommendation is the representation of a recommendation
type Recommendation struct {
	ID         int
	IdPosition *Position
	IdUser     *User
	IdUser2    *User
}

//GetRecommendations return all recommendations from the database
func (db *DB) GetRecommendations() ([]*Recommendation, error) {

	rows, err := db.Query("SELECT * FROM recommendations")
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
		err = rows.Scan(&rec.ID, &idPos, &idUsr, &idUsr2)
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

		recs = append(recs, rec)

	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return recs, nil
}

//GetRecommendation returns a single recommendation based on it's id
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

		err = rows.Scan(&rec.ID, &idPos, &idUsr, &idUsr2)
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
