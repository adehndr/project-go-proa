package helper

import (
	"encoding/json"
	"net/http"

	"example.com/adehndr/project_go_proa/model/web"
)

func FetchTasks() (web.WebResponseRequest, error) {
	var err error
	var client = &http.Client{}
	var tempWebResponse web.WebResponseRequest = web.WebResponseRequest{}
	request, err := http.NewRequest("GET", "http://localhost:3000/tasks", nil)
	if err != nil {
		return tempWebResponse, err
	}
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
