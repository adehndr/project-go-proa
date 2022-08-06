package service

import (
	"context"
	"fmt"
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

func (service *TaskListServiceImplementation) FindAll(ctx context.Context) ([]web.TaskResponse, error) {
	data, err := service.taskListRepository.FindAll(ctx)
	var webTaskResponseList []web.TaskResponse = []web.TaskResponse{}
	if err != nil {
		fmt.Println("Error di service", err.Error())
		log.Fatal(err)
		return webTaskResponseList, err
	}
	for _, te := range data {
		webTaskResponse := web.TaskResponse{
			Id:         te.Id,
			TaskDetail: te.TaskDetail,
			Asignee:    te.Asignee,
			Deadline:   te.Deadline,
			IsFinished: te.IsFinished,
		}
		webTaskResponseList = append(webTaskResponseList, webTaskResponse)
	}
	return webTaskResponseList, nil
}

func (service *TaskListServiceImplementation) FindById(ctx context.Context, id int) (web.TaskResponse, error) {
	var taskResponse web.TaskResponse = web.TaskResponse{}
	res, err := service.taskListRepository.FindById(ctx, id)
	if err != nil {
		// Check the error type
		return taskResponse, err
	}
	taskResponse.Id = res.Id
	taskResponse.TaskDetail = res.TaskDetail
	taskResponse.Asignee = res.Asignee
	taskResponse.Deadline = res.Deadline
	taskResponse.IsFinished = res.IsFinished
	return taskResponse, nil
}

/*
	Insert row, after successfull, it will return the new object saved on database with its id
*/
func (service *TaskListServiceImplementation) Create(ctx context.Context, webTaskCreateRequest web.TaskCreateRequest) (web.TaskResponse, error) {
	var response web.TaskResponse = web.TaskResponse{}
	taskObject := entity.TaskEntity{
		TaskDetail: webTaskCreateRequest.TaskDetail,
		Asignee:    webTaskCreateRequest.Asignee,
		Deadline:   webTaskCreateRequest.Deadline,
		IsFinished: webTaskCreateRequest.IsFinished,
	}
	data, err := service.taskListRepository.Create(ctx, taskObject)
	if err != nil {
		return response, err
	}
	response.Id = data.Id
	response.TaskDetail = data.TaskDetail
	response.Asignee = data.Asignee
	response.Deadline = data.Deadline
	response.IsFinished = data.IsFinished
	return response, nil
}

func (service *TaskListServiceImplementation) Update(ctx context.Context, task web.TaskUpdateRequest) (web.TaskResponse, error) {
	var webResponse web.TaskResponse = web.TaskResponse{}
	taskFounded, err := service.taskListRepository.FindById(ctx, task.Id)
	if err != nil {
		return webResponse, err
	}
	taskFounded.TaskDetail = task.TaskDetail
	taskFounded.Asignee = task.Asignee
	taskFounded.Deadline = task.Deadline
	taskFounded.IsFinished = task.IsFinished
	taskUpdated, err := service.taskListRepository.Update(ctx, taskFounded)
	if err != nil {
		return webResponse, err
	}
	webResponse.Id = taskUpdated.Id
	webResponse.TaskDetail = taskUpdated.TaskDetail
	webResponse.Deadline = taskUpdated.Deadline
	webResponse.IsFinished = taskUpdated.IsFinished
	return webResponse, nil
}

func (service *TaskListServiceImplementation) Delete(ctx context.Context, id int) error {
	err := service.taskListRepository.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
