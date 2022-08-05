package repository

import (
	"context"
	
	"example.com/adehndr/project_go_proa/entity"
)

type AnimeListRepository interface {
	FindAll(ctx context.Context) ([]entity.AnimeEntity, error)
}
