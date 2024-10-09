package database

import (
	"github.com/jmoiron/sqlx"
	"time"
)

var database *sqlx.DB

func InitializeDatabase() error {

	db, error := sqlx.Connect("mysql", "user=root dbname=anankes sslmode=disable")

	if error != nil {
		return error
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(time.Minute * 3)

	database = db

	return nil
}

func GetConnection() *sqlx.DB {
	return database
}
