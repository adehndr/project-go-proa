package service

import (
	"context"

	"example.com/adehndr/project_go_proa/model/web"
)

type TaskListService interface {
	FindAll(ctx context.Context) ([]web.TaskResponse, error)
	FindById(ctx context.Context, id int) (web.TaskResponse, error)
	Create(ctx context.Context, webTaskCreateRequest web.TaskCreateRequest) (web.TaskResponse, error)
	Update(ctx context.Context, task web.TaskUpdateRequest) (web.TaskResponse, error)
	Delete(ctx context.Context, id int) error
}
