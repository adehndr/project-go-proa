package repository

import (
	"context"
	"database/sql"
	"log"

	"github.com/adehndr/anime-databases/entity"
)

type AnimeListRepositoryImpl struct {
	DB *sql.DB
}

func NewAnimeListRepository(db *sql.DB) AnimeListRepository {
	return &AnimeListRepositoryImpl{
		DB: db,
	}
}

func(repository *AnimeListRepositoryImpl) FindAll(ctx context.Context) ([]entity.AnimeEntity, error) {
	sqlQuery := "SELECT id,title,description, episodes, aired,finished from anime_list"
	rows , err := repository.DB.QueryContext(ctx,sqlQuery)
	if err != nil {
		log.Fatal(err)
		return nil,err
	}
	defer rows.Close()
	var listAnimeList []entity.AnimeEntity
	for rows.Next() {
		objAnimeList := entity.AnimeEntity{}
		err := rows.Scan(
			&objAnimeList.Id,
			&objAnimeList.Title,
			&objAnimeList.Description,
			&objAnimeList.Episodes,
			&objAnimeList.Aired,
			&objAnimeList.Finished,
		)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		listAnimeList = append(listAnimeList, objAnimeList)
	}
	return listAnimeList,nil
}