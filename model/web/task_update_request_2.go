package web

type TaskUpdateRequest2 struct {
	Id         int    `json:"id"`
	TaskDetail string `json:"task_detail"`
	Asignee    string `json:"assignee"`
	Deadline   string `json:"deadline"`
	IsFinished bool   `json:"is_finished"`
}
