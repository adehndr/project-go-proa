package repository

import (
	"context"

	"github.com/adehndr/anime-databases/entity"
)

type AnimeListRepository interface {
	FindAll(ctx context.Context) ([]entity.AnimeEntity, error)
}
