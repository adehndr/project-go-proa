package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"example.com/adehndr/project_go_proa/app"
	"example.com/adehndr/project_go_proa/controller"
	"example.com/adehndr/project_go_proa/repository"
	"example.com/adehndr/project_go_proa/service"
	"github.com/julienschmidt/httprouter"
)

var (
	dbMySql            *sql.DB                       = app.OpenDatabaseConnection()
	taskListRepository repository.TaskListRepository = repository.NewTaskListRepository(dbMySql)
	taskListSevice     service.TaskListService       = service.NewTaskListService(taskListRepository)
	taskController     controller.TaskController     = controller.NewTaskController(taskListSevice)
)

func main() {

	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}
	defer dbMySql.Close()

	err := dbMySql.Ping()
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	tesHandler := httprouter.New()
	tesHandler.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {fmt.Fprint(w,"Success")})
	server := http.Server{
		Addr:    ":" + port,
		Handler: tesHandler,
	}
	log.Fatal(server.ListenAndServe())
}
