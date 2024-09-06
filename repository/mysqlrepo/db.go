package mysqlrepo

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type MYSQlDB struct {
	db *sql.DB
}

func New() *MYSQlDB {
	db, err := sql.Open("mysql", "gameapp:gameappt0lk2o20@(localhost:3306)/gameapp_db")
	if err != nil {
		panic(fmt.Errorf("can't open mysql db: %v", err))
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return &MYSQlDB{
		db: db,
	}
}
