package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"example.com/adehndr/project_go_proa/model/web"
	"example.com/adehndr/project_go_proa/service"
	"github.com/julienschmidt/httprouter"
)

type TaskControllerImpl struct {
	Service service.TaskListService
}

func NewTaskController(service service.TaskListService) TaskController {
	return &TaskControllerImpl{
		Service: service,
	}
}

func (controller *TaskControllerImpl) FindAll(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var webResponse web.WebResponse = web.WebResponse{}
	encoder := json.NewEncoder(w)
	response, err := controller.Service.FindAll(r.Context())
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	if err != nil {
		webResponse.Message = err.Error()
		webResponse.Status = "fail"
		w.WriteHeader(http.StatusBadRequest)
		log.Fatal(err)
	} else {
		webResponse.Status = "success"
		webResponse.Data = response
		w.WriteHeader(http.StatusOK)
	}
	encoder.Encode(webResponse)
}

func (controller *TaskControllerImpl) FindById(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	idParams, err := strconv.Atoi(params.ByName("taskid"))
	encoder := json.NewEncoder(w)
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	if err != nil {
		log.Fatal(err)
	}
	result, err := controller.Service.FindById(r.Context(), idParams)
	webResponse := web.WebResponse{}
	if err != nil {
		webResponse.Message = err.Error()
		w.WriteHeader(http.StatusBadRequest)
	} else {
		webResponse.Status = "success"
		webResponse.Data = result
		w.WriteHeader(http.StatusOK)
	}
	encoder.Encode(webResponse)
}

func (controller *TaskControllerImpl) Create(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var taskCreateRequest web.TaskCreateRequest = web.TaskCreateRequest{}
	webResponse := web.WebResponse{}
	decoder := json.NewDecoder(r.Body)
	encoder := json.NewEncoder(w)
	err := decoder.Decode(&taskCreateRequest)
	if err != nil {
		panic(err)
	}
	result, err := controller.Service.Create(r.Context(), taskCreateRequest)
	if err != nil {
		w.WriteHeader(http.StatusCreated)
		panic(err)
	} else {
		webResponse.Status = "success"
		webResponse.Data = result
		w.WriteHeader(http.StatusOK)
	}
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	encoder.Encode(webResponse)
}
func (controller *TaskControllerImpl) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	idParams, err := strconv.Atoi(params.ByName("taskid"))
	var taskUpdateRequest web.TaskUpdateRequest = web.TaskUpdateRequest{}
	var webResponse web.WebResponse = web.WebResponse{}
	decoder := json.NewDecoder(r.Body)
	if err != nil {
		panic(err)
	}
	err = decoder.Decode(&taskUpdateRequest)
	taskUpdateRequest.Id = idParams
	result, err := controller.Service.Update(r.Context(), taskUpdateRequest)
	if err != nil {
		panic(err)
	}
	encoder := json.NewEncoder(w)
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	webResponse.Data = result
	encoder.Encode(webResponse)
}
func (controller *TaskControllerImpl) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	idParams, err := strconv.Atoi(params.ByName("taskid"))
	var webResponse web.WebResponse = web.WebResponse{}
	if err != nil {
		panic(err)
	}
	encoder := json.NewEncoder(w)
	err = controller.Service.Delete(r.Context(),idParams)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}else {
		webResponse.Status = "success"
		w.WriteHeader(http.StatusOK)
	}
	encoder.Encode(webResponse)
}
