package service

import (
	"context"

	"example.com/adehndr/project_go_proa/model/web"
)

type AnimeListService interface {
	FindAll(ctx context.Context) (web.WebResponse, error)
}