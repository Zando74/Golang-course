package model

import "todo-app/internal/app/domain/entity"

type Task struct {
	Id          string `gorm:"primaryKey;type:uuid"`
	Title       string `gorm:"column:title;type:text"`
	Description string `gorm:"column:description;type:text"`
	Status      int    `gorm:"column:status;type:int"`
}

func (Task) TableName() string {
	return "task"
}

func NewTaskModelFromEntity(task entity.Task) Task {
	return Task{
		Id:          task.Id,
		Title:       task.Title.String(),
		Description: task.Description.String(),
		Status:      int(task.Status.Get()),
	}
}
