package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"example.com/adehndr/project_go_proa/app"
	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
)

func main() {

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	port := os.Getenv("PORT")

	psqlInfo := fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s sslmode=require",
		dbHost, dbUser, dbPassword, dbPort, dbName)

	db := app.OpenDatabaseConnection(psqlInfo)

	if port == "" {
		port = "3000"
	}
	defer db.Close()
	tesHandler := httprouter.New()
	tesHandler.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) { fmt.Fprint(w, "Success") })
	tesHandler.GET("/test", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		var tempStr []string = []string{}
		data, err := db.QueryContext(r.Context(), "SELECT task_detail FROM task_table")
		if err != nil {
			panic(err)
		}
		for data.Next() {
			var tempData string
			err = data.Scan(&tempData)
			if err != nil {
				panic(err)
			}
			tempStr = append(tempStr, tempData)
		}
		fmt.Fprint(w, tempStr)
	})
	server := http.Server{
		Addr:    ":" + port,
		Handler: tesHandler,
	}
	log.Fatal(server.ListenAndServe())
}
