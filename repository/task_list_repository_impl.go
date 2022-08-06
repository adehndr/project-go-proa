package repository

import (
	"context"
	"database/sql"
	"errors"
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
	var entityTask entity.TaskEntity = entity.TaskEntity{}
	querySQL := "SELECT id,task_detail,assignee,deadline,is_finished from task_table where id = $1"
	res, err := repository.DB.QueryContext(ctx, querySQL, id)
	if err != nil {
		log.Fatal(err)
		return entity.TaskEntity{}, err
	}
	if res.Next() {
		err := res.Scan(
			&entityTask.Id,
			&entityTask.TaskDetail,
			&entityTask.Asignee,
			&entityTask.Deadline,
			&entityTask.IsFinished,
		)
		if err != nil {
			log.Fatal(err)
			return entityTask, err
		}
		return entityTask, nil
	} else {
		return entityTask, errors.New("Task Not Found")
	}
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
	querySQL := "UPDATE task_table SET task_detail = $1, assignee = $2, deadline = $3, is_finished = $4 where id = $5 returning id"
	_, err := repository.DB.ExecContext(ctx, querySQL, task.TaskDetail, task.Asignee, task.Deadline, task.IsFinished, task.Id)
	if err != nil {
		return entity.TaskEntity{}, err
	}
	return task, nil
}

func (repository *TaskListRepositoryImpl) Delete(ctx context.Context, id int) error {
	querySQL := "DELETE FROM task_table where id = $1"
	_, err := repository.DB.ExecContext(ctx, querySQL, id)
	if err != nil {
		return err
	}
	return nil
}
