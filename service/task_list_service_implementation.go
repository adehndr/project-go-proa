package service

import (
	"context"
	"log"

	"example.com/adehndr/project_go_proa/entity"
	"example.com/adehndr/project_go_proa/model/web"
	"example.com/adehndr/project_go_proa/repository"
)

type TaskListServiceImplementation struct {
	taskListRepository repository.TaskListRepository
}

func NewTaskListService(taskListRepository repository.TaskListRepository) TaskListService {
	return &TaskListServiceImplementation{
		taskListRepository: taskListRepository,
	}
}

func (service *TaskListServiceImplementation) FindAll(ctx context.Context) (web.WebResponse, error) {
	var response web.WebResponse = web.WebResponse{}
	data, err := service.taskListRepository.FindAll(ctx)
	if err != nil {
		log.Fatal(err)
		return response, err
	}
	response.Code = 200
	response.Status = "Success"
	response.Data = data
	return response, nil
}

func (service *TaskListServiceImplementation) FindById(ctx context.Context, id int) (web.WebResponse, error) {
	var webResponse web.WebResponse
	return webResponse, nil
}
/* 
	Insert row, after successfull, it will return the new object saved on database with its id
*/
func (service *TaskListServiceImplementation) Create(ctx context.Context, webTaskCreateRequest web.TaskCreateRequest) (web.WebResponse, error) {
	var response web.WebResponse = web.WebResponse{}
	taskObject := entity.TaskEntity{
		TaskDetail: webTaskCreateRequest.TaskDetail,
		Asignee:    webTaskCreateRequest.Asignee,
		Deadline:   webTaskCreateRequest.Deadline,
		IsFinished: webTaskCreateRequest.IsFinished,
	}
	data, err := service.taskListRepository.Create(ctx, taskObject)
	taskObject.Id = data.Id
	if err != nil {
		response.Code = 400
		response.Status = err.Error()
		response.Data = nil
		return response, err
	}
	response.Code = 200
	response.Status = "Success"
	response.Data = taskObject
	return response, nil
}

func (service *TaskListServiceImplementation) Update(ctx context.Context, task entity.TaskEntity) (web.WebResponse, error) {
	var webResponse web.WebResponse
	return webResponse, nil
}

func (service *TaskListServiceImplementation) Delete(ctx context.Context, id int) (web.WebResponse, error) {
	var webResponse web.WebResponse
	return webResponse, nil
}
