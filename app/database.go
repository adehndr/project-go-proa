package app

import (
	"database/sql"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
)

const (
	host     = "postgres://ibfvfqhqbzluor:c41029e235d47f47a2af5f64928b6000165a5e36b8e58ed5631fbf87fd303648@ec2-54-225-234-165.compute-1.amazonaws.com:5432/d4nk1d6613bi4t"
	port     = 5432
	user     = "postgres"
	password = "root"
	dbname   = "postgres"
)

func OpenDatabaseConnection() *sql.DB {
	/* 	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
	"password=%s dbname=%s sslmode=disable",
	host, port, user, password, dbname) */
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	db.SetConnMaxIdleTime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	db.Ping()
	return db
}
