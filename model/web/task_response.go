package web

import "time"

type TaskResponse struct {
	Id         int       `json:"id"`
	TaskDetail string    `json:"task_detail"`
	Asignee    string    `json:"assignee"`
	Deadline   time.Time `json:"deadline"`
	IsFinished bool      `json:"is_finished"`
}