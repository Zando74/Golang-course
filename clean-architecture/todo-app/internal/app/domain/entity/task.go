package entity

import (
	"todo-app/internal/app/domain/valueobject"
)

type Task struct {
	Id          string                      `json:"id"`
	Title       valueobject.TaskTitle       `json:"title"`
	Description valueobject.TaskDescription `json:"description"`
	Status      valueobject.TaskState       `json:"status"`
}
