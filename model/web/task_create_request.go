package web

import "time"

type TaskCreateRequest struct {
	TaskDetail string    `json:"task_detail"`
	Asignee    string    `json:"assignee"`
	Deadline   time.Time `json:"deadline"`
	IsFinished bool      `json:"is_finished"`
}