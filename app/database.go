package app

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/lib/pq"
)

func OpenDatabaseConnection(psqlInfo string) *sql.DB {
	log.Fatal(psqlInfo)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	db.SetConnMaxIdleTime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	db.Ping()
	return db
}
