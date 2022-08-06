package service

import (
	"context"

	"example.com/adehndr/project_go_proa/entity"
	"example.com/adehndr/project_go_proa/model/web"
)

type TaskListService interface {
	FindAll(ctx context.Context) (web.WebResponse, error)
	FindById(ctx context.Context, id int) (web.WebResponse,error)
	Create(ctx context.Context, webTaskCreateRequest web.TaskCreateRequest) (web.WebResponse,error)
	Update(ctx context.Context, task entity.TaskEntity) (web.WebResponse,error)
	Delete(ctx context.Context, id int) (web.WebResponse,error)
}