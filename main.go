package main

import (
	"database/sql"
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
	taskRouter         *httprouter.Router            = (*httprouter.Router)(app.NewRouter(taskController))
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
	server := http.Server{
		Addr:    ":" + port,
		Handler: taskRouter,
	}
	log.Fatal(server.ListenAndServe())
}
