package helper

import (
	"encoding/json"
	"net/http"
	"sort"

	"example.com/adehndr/project_go_proa/model/web"
)

func FetchTasks() (web.WebResponseRequest, error) {
	var err error
	var client = &http.Client{}
	var tempWebResponse web.WebResponseRequest = web.WebResponseRequest{}
	request, err := http.NewRequest("GET", "https://go-project-proa.herokuapp.com/api/tasks", nil)
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
	var tempArr []web.TaskResponse = []web.TaskResponse{}
	if tempWebResponse.Status == "success" && len(tempWebResponse.Data) > 0 {
		tempArr = append(tempArr, tempWebResponse.Data...)
		sort.Slice(tempArr, func(i, j int) bool {
			return tempArr[i].Id < tempArr[j].Id
		})
		tempWebResponse.Data = tempArr
	}
	return tempWebResponse, nil
}
