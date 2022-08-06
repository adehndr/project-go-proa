package main

import (
	"context"
	"fmt"
	"log"

	"example.com/adehndr/project_go_proa/app"
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
	/* 		fmt.Println(taskListSevice.Create(ctx, web.TaskCreateRequest{
	   		TaskDetail: "Add feature",
	   		Deadline:   time.Now(),
	   		Asignee:    "Ade",
	   		IsFinished: false,
	   	})) */
	/* 	fmt.Println(
		taskListSevice.Update(
			ctx,
			web.TaskUpdateRequest{
				Id:         9,
				TaskDetail: "Fix ",
				Asignee:    "Hendra",
				Deadline:   time.Now(),
				IsFinished: false,
			},
		),
	) */
	fmt.Println(taskListSevice.FindAll(ctx))
	// fmt.Println(taskListSevice.FindById(ctx, 21))
	// fmt.Println(taskListSevice.Delete(ctx, 11))
}
