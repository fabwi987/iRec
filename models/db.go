package models

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type Datastore interface {
	GetUsers() ([]*User, error)
	GetUser(userid int) (*User, error)
	GetPositions() ([]*Position, error)
	GetPosition(positionid int) (*Position, error)
	GetRecommendations() ([]*Recommendation, error)
	GetRecommendation(recommendationid int) (*Recommendation, error)
}

type DB struct {
	*sql.DB
}

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
