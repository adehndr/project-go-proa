package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"example.com/adehndr/project_go_proa/app"
	"example.com/adehndr/project_go_proa/controller"
	"example.com/adehndr/project_go_proa/helper"
	"example.com/adehndr/project_go_proa/repository"
	"example.com/adehndr/project_go_proa/service"
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

	db, err := sql.Open("postgres", psqlInfo)
	db.SetConnMaxIdleTime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	helper.InitializeDatabase(db,context.Background())
	if err != nil {
		panic(err)
	}
	if port == "" {
		port = "3000"
	}
	defer db.Close()
	taskRepository := repository.NewTaskListRepository(db)
	taskService := service.NewTaskListService(taskRepository)
	taskController := controller.NewTaskController(taskService)
	router := app.NewRouter(taskController)
	server := http.Server{
		Addr:    ":" + port,
		Handler: router,
	}
	log.Fatal(server.ListenAndServe())
}
