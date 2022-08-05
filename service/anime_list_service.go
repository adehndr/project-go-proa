package service

import (
	"context"

	"github.com/adehndr/anime-databases/model/web"
)

type AnimeListService interface {
	FindAll(ctx context.Context) (web.WebResponse, error)
}