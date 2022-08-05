package service

import (
	"context"
	"log"

	"example.com/adehndr/project_go_proa/model/web"
	"example.com/adehndr/project_go_proa/repository"
)

type AnimeListServiceImplementation struct {
	animeListRepository repository.AnimeListRepository
}

func NewAnimeListService(animeListRepository repository.AnimeListRepository) AnimeListService {
	return &AnimeListServiceImplementation{
		animeListRepository: animeListRepository,
	}
}

func (service *AnimeListServiceImplementation) FindAll(ctx context.Context) (web.WebResponse, error) {
	var response web.WebResponse = web.WebResponse{}
	data, err := service.animeListRepository.FindAll(ctx)
	if err != nil {
		log.Fatal(err)
		return response, err
	}
	response.Code = 200
	response.Status = "Success"
	response.Data = data
	return response, nil
}
