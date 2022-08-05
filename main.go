package main

import (
	"context"
	"fmt"
	"log"

	"github.com/adehndr/anime-databases/app"
	"github.com/adehndr/anime-databases/repository"
	"github.com/adehndr/anime-databases/service"
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
