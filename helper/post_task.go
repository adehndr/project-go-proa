package helper

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"example.com/adehndr/project_go_proa/model/web"
)

func PostTask(r *http.Request) (web.TaskCreateResponse, error) {
	fmt.Println("Ini create")
	var err error
	var payloadRequest web.TaskCreateRequest = web.TaskCreateRequest{}
	payloadRequest.TaskDetail = r.FormValue("detail_task")
	payloadRequest.Asignee = r.FormValue("assignee")
	tempVar := r.FormValue("deadline")
	myDate, err := time.Parse("2006-01-02", tempVar)
	payloadRequest.Deadline = myDate
	payloadRequest.IsFinished = false

	finalPayload, err := json.Marshal(payloadRequest)
	if err != nil {
		panic(err)
	}
	request, err := http.NewRequest("POST", "https://go-project-proa.herokuapp.com/api/tasks", bytes.NewBuffer(finalPayload))
	request.Header.Set("Content-Type", "application/json")

	var tempWebResponse web.TaskCreateResponse = web.TaskCreateResponse{}
	if err != nil {
		return tempWebResponse, err
	}

	var client = &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return tempWebResponse, err
	}
	defer response.Body.Close()
	decoder := json.NewDecoder(response.Body)
	err = decoder.Decode(&tempWebResponse)
	if err != nil {
		return tempWebResponse, err
	}
	return tempWebResponse, nil
}
