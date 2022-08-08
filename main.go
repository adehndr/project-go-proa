package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

/* var (
	dbMySql            *sql.DB                       = app.OpenDatabaseConnection(dbUser, dbPassword, dbHost, dbPort, dbName)
	taskListRepository repository.TaskListRepository = repository.NewTaskListRepository(dbMySql)
	taskListSevice     service.TaskListService       = service.NewTaskListService(taskListRepository)
	taskController     controller.TaskController     = controller.NewTaskController(taskListSevice)
) */

func main() {

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	port := os.Getenv("PORT")

	psqlInfo := fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s sslmode=require",
		dbHost, dbUser, dbPassword, dbPort, dbName)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}

	if port == "" {
		port = "3000"
	}
	defer db.Close()
	tesHandler := httprouter.New()
	tesHandler.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) { fmt.Fprint(w, "Success") })
	server := http.Server{
		Addr:    ":" + port,
		Handler: tesHandler,
	}
	log.Fatal(server.ListenAndServe())
}
