package sql

import (
	"database/sql"
	"fmt"
	"time"
)

type MYSQlDB struct {
	db *sql.DB
}

func New() *MYSQlDB {
	db, err := sql.Open("mysql", "gameapp:gameappt0lk2o20@(localhost:3306)/gameapp_db")
	if err != nil {
		panic(fmt.Errorf("can't open sql db: %v", err))
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return &MYSQlDB{
		db: db,
	}
}
