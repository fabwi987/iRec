package models

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

//Datastore gathers the methods available for the environment
type Datastore interface {
	GetUsers() ([]*User, error)
	GetUser(userid int) (*User, error)
	GetPositions() ([]*Position, error)
	GetPosition(positionid int) (*Position, error)
	GetRecommendations() ([]*Recommendation, error)
	GetRecommendation(recommendationid int) (*Recommendation, error)
}

//DB hold the db connection
type DB struct {
	*sql.DB
}

//NewDatabase initiates the database connections
func NewDatabase(dataSourceName string) (*DB, error) {

	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}
	return &DB{db}, nil
}
