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
	dbMySql := app.OpenDatabaseConnection()
	defer dbMySql.Close()

	err := dbMySql.Ping()
	if err != nil {
		log.Fatal(err)
	}
	animeListRepository := repository.NewAnimeListRepository(dbMySql)
	animeListSevice := service.NewAnimeListService(animeListRepository)
	fmt.Println(animeListSevice.FindAll(context.Background()))
}
