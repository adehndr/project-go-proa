package main

import (
	"log"
	"net/http"
	"os"

	"example.com/adehndr/project_go_proa/app"
	"example.com/adehndr/project_go_proa/controller"
	"example.com/adehndr/project_go_proa/repository"
	"example.com/adehndr/project_go_proa/service"
)

func main() {

	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}

	dbMySql := app.OpenDatabaseConnection()
	defer dbMySql.Close()

	err := dbMySql.Ping()
	if err != nil {
		log.Fatal(err)
	}
	taskListRepository := repository.NewTaskListRepository(dbMySql)
	taskListSevice := service.NewTaskListService(taskListRepository)
	taskController := controller.NewTaskController(taskListSevice)
	taskRouter := app.NewRouter(taskController)

	server := http.Server{
		Addr:    ":" + port,
		Handler: taskRouter,
	}
	log.Fatal(server.ListenAndServe())
}
