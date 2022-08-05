package app

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func OpenDatabaseConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/pustaka_api")
	if err != nil {
		log.Fatal(err)
	}
	db.SetConnMaxIdleTime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	db.Ping()
	return db
}
