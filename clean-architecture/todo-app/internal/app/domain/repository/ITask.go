package repository

import "todo-app/internal/app/domain/entity"

type ITask interface {
	Create(entity.Task) (*entity.Task, error)
	Get(id string) (*entity.Task, error)
	Index() ([]*entity.Task, error)
	Update(task entity.Task) (*entity.Task, error)
	Delete(task entity.Task) (*entity.Task, error)
}
