package entity

import (
	"time"
)

type TaskEntity struct {
	Id         int
	TaskDetail string
	Asignee    string
	Deadline   time.Time
	IsFinished bool
}
