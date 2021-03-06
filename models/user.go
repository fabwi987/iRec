package models

import _ "github.com/go-sql-driver/mysql"

//User is the representation of a user
type User struct {
	Id    int
	Type  int
	Name  string
	Mail  string
	Phone string
}

//GetUsers return all users from the database
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

//GetUser returns a single user based on it's id
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

//CreateUser creates a new user inte the database
func (db *DB) CreateUser(NewUser *User) (int64, error) {

	stmt, err := db.Prepare("INSERT users SET usertype=?,name=?,mail=?,phone=?")
	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec(NewUser.Type, NewUser.Name, NewUser.Mail, NewUser.Phone)
	if err != nil {
		return 0, err
	}
	LastID, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return LastID, nil
}
