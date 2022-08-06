package repository

import (
	"context"
	"database/sql"
	"log"

	"example.com/adehndr/project_go_proa/entity"
)

type TaskListRepositoryImpl struct {
	DB *sql.DB
}

func NewTaskListRepository(db *sql.DB) TaskListRepository {
	return &TaskListRepositoryImpl{
		DB: db,
	}
}

func (repository *TaskListRepositoryImpl) FindAll(ctx context.Context) ([]entity.TaskEntity, error) {
	sqlQuery := "SELECT id,task_detail,assignee, deadline, is_finished from task_table"
	rows, err := repository.DB.QueryContext(ctx, sqlQuery)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer rows.Close()
	var taskList []entity.TaskEntity
	for rows.Next() {
		objTask := entity.TaskEntity{}
		err := rows.Scan(
			&objTask.Id,
			&objTask.TaskDetail,
			&objTask.Asignee,
			&objTask.Deadline,
			&objTask.IsFinished,
		)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		taskList = append(taskList, objTask)
	}
	return taskList, nil
}

func (repository *TaskListRepositoryImpl) FindById(ctx context.Context, id int) (entity.TaskEntity, error) {
	var taskEntity entity.TaskEntity
	return taskEntity, nil
}

/* 
	Return the same Task, but add the id got from the inserted row on database
*/
func (repository *TaskListRepositoryImpl) Create(ctx context.Context, task entity.TaskEntity) (entity.TaskEntity, error) {
	var lastInsertedId int
	/*
		It should use ExecContext for insert,update, and delete
		but because I need the id from the inserted row, so it use QueryRowContext
	*/
	result := repository.DB.QueryRowContext(ctx, "INSERT INTO task_table(task_detail,assignee,deadline,is_finished) values($1,$2,$3,$4) RETURNING id", task.TaskDetail, task.Asignee, task.Deadline, task.IsFinished)
	result.Scan(&lastInsertedId)
	task.Id = lastInsertedId
	return task, nil
}

func (repository *TaskListRepositoryImpl) Update(ctx context.Context, task entity.TaskEntity) (entity.TaskEntity, error) {
	var taskEntity entity.TaskEntity
	return taskEntity, nil
}

func (repository *TaskListRepositoryImpl) Delete(ctx context.Context, id int) error {
	return nil
}
