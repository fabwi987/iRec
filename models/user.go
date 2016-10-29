package models

import _ "github.com/go-sql-driver/mysql"

type User struct {
	Id    int
	Type  int
	Name  string
	Mail  string
	Phone string
}

func (db *DB) GetUsers() ([]*User, error) {

	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	usrs := make([]*User, 0)
	for rows.Next() {
		usr := new(User)
		err = rows.Scan(&usr.Id, &usr.Type, &usr.Name, &usr.Mail, &usr.Phone)
		if err != nil {
			return nil, err
		}
		usrs = append(usrs, usr)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return usrs, nil
}

func (db *DB) GetUser(userid int) (*User, error) {

	stmt, err := db.Prepare("SELECT * FROM users WHERE idusers = ?")
	defer stmt.Close()
	rows, err := stmt.Query(userid)
	defer rows.Close()
	us := new(User)

	for rows.Next() {
		err = rows.Scan(&us.Id, &us.Type, &us.Name, &us.Mail, &us.Phone)
		if err != nil {
			return nil, err
		}
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return us, nil

}
