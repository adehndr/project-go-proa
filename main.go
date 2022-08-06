package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"example.com/adehndr/project_go_proa/app"
	"example.com/adehndr/project_go_proa/model/web"
	"example.com/adehndr/project_go_proa/repository"
	"example.com/adehndr/project_go_proa/service"
)

func main() {
	ctx := context.Background()
	dbMySql := app.OpenDatabaseConnection()
	defer dbMySql.Close()

	err := dbMySql.Ping()
	if err != nil {
		log.Fatal(err)
	}
	taskListRepository := repository.NewTaskListRepository(dbMySql)
	taskListSevice := service.NewTaskListService(taskListRepository)
	fmt.Println(taskListSevice.Create(ctx, web.TaskCreateRequest{
		TaskDetail: "Fix bug 2",
		Deadline:   time.Now(),
		Asignee:    "Setiawan 2",
		IsFinished: false,
	}))
	fmt.Println(taskListSevice.FindAll(ctx))
}
