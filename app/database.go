package app

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

func OpenDatabaseConnection(dbUser string, dbPassword string, dbHost string, dbPort string, dbName string) *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s sslmode=require",
		dbHost, dbUser, dbPassword, dbPort, dbName)
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
