package controller

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"example.com/adehndr/project_go_proa/helper"
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

func (controller *TaskControllerImpl) Home(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	if r.Method != http.MethodPost {
		t := template.New("index.gohtml")
		data, err := helper.FetchTasks()
		if err != nil {
			panic(err)
		}
		t, err = template.ParseGlob("./templates/*.gohtml")
		if err != nil {
			panic(err)
		}
		t.ExecuteTemplate(w, "index.gohtml", map[string]interface{}{
			"DetailTask": data.Data,
		})
		return
	}

}

func (controller *TaskControllerImpl) Detail(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var err error
	query := r.URL.Query()
	action := query.Get("action")
	idParams := query.Get("id")
	t := template.Must(template.ParseGlob("./templates/*.gohtml"))
	if r.Method != http.MethodPost {
		var resultResponse web.TaskUpdateRequest2
		if action != "" {
			idTaskInt, err := strconv.Atoi(idParams)
			if err != nil {
				panic(err)
			}
			result, err := controller.Service.FindById(r.Context(), idTaskInt)
			if err != nil {
				panic(err)
			}
			resultResponse = web.TaskUpdateRequest2{
				Id:         result.Id,
				Asignee:    result.Asignee,
				TaskDetail: result.TaskDetail,
				IsFinished: result.IsFinished,
				Deadline:   result.Deadline.Format("2006-01-02"),
			}
			t.ExecuteTemplate(w, "create.gohtml", map[string]interface{}{
				"DetailTask": resultResponse,
			})
		} else {
			t.ExecuteTemplate(w, "create.gohtml", map[string]interface{}{
				"DetailTask": "",
			})
		}

		return
	}
	action = r.FormValue("action")
	if action == "create" {
		_, err = helper.PostTask(r)
	}else {
		_, err = helper.UpdateTask(r)
	}

	if err != nil {
		panic(err)
	}
	http.Redirect(w, r, "https://go-project-proa.herokuapp.com/", 301)
}

func (controller *TaskControllerImpl) Change(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	idParams, err := strconv.Atoi(params.ByName("taskid"))
	action := r.URL.Query().Get("action")
	findResponse, err := controller.Service.FindById(r.Context(), idParams)
	if err != nil {
		panic(err)
	}
	if action == "done" {
		webTaskReq := web.TaskUpdateRequest{
			Id:         findResponse.Id,
			TaskDetail: findResponse.TaskDetail,
			Asignee:    findResponse.Asignee,
			Deadline:   findResponse.Deadline,
			IsFinished: true,
		}
		controller.Service.Update(r.Context(), webTaskReq)
		http.Redirect(w, r, "https://go-project-proa.herokuapp.com/", 301)
	} else {
		http.Redirect(w, r, "https://go-project-proa.herokuapp.com/detail", 301)
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
	encoder := json.NewEncoder(w)
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	if err != nil {
		log.Fatal(err)
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
	err = controller.Service.Delete(r.Context(), idParams)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		webResponse.Status = "success"
		w.WriteHeader(http.StatusOK)
	}
	encoder.Encode(webResponse)
}
