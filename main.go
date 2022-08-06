package main

import (
	"log"
	"net/http"

	"example.com/adehndr/project_go_proa/app"
	"example.com/adehndr/project_go_proa/controller"
	"example.com/adehndr/project_go_proa/repository"
	"example.com/adehndr/project_go_proa/service"
)

func main() {
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
		Addr:    "localhost:3000",
		Handler: taskRouter,
	}
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
