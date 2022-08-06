package repository

import (
	"context"

	"example.com/adehndr/project_go_proa/entity"
)

type TaskListRepository interface {
	FindAll(ctx context.Context) ([]entity.TaskEntity, error)
	FindById(ctx context.Context, id int) (entity.TaskEntity,error)
	Create(ctx context.Context, task entity.TaskEntity) (entity.TaskEntity,error)
	Update(ctx context.Context, task entity.TaskEntity) (entity.TaskEntity,error)
	Delete(ctx context.Context, id int) error
}
