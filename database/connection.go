package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

var database *sql.DB

func InitializeDatabase() error {

	db, error := sql.Open("mysql", "root@/anankes")

	if error != nil {
		return error
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(10)
	db.SetConnMaxLifetime(time.Minute * 3)

	database = db

	return nil
}

func GetConnection() *sql.DB {
	return database
}
