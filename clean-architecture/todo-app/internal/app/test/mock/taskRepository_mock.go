package mock

import (
	"todo-app/internal/app/domain/entity"

	"golang.org/x/exp/maps"
)

type UnexistingTaskError struct{}

func (e *UnexistingTaskError) Error() string {
	return "Unexisting Task"
}

type MTaskRepository struct {
	tasks map[string]*entity.Task
}

func NewMTaskRepository() *MTaskRepository {
	return &MTaskRepository{tasks: make(map[string]*entity.Task)}
}

func (mock *MTaskRepository) Create(task entity.Task) (*entity.Task, error) {
	mock.tasks[task.Id] = &task
	return &task, nil
}

func (mock *MTaskRepository) Get(id string) (*entity.Task, error) {
	existingTask := mock.tasks[id]

	if existingTask == nil {
		return nil, &UnexistingTaskError{}
	}

	return existingTask, nil
}

func (mock *MTaskRepository) Index() ([]*entity.Task, error) {
	return maps.Values(mock.tasks), nil
}

func (mock *MTaskRepository) Update(task entity.Task) (*entity.Task, error) {
	mock.tasks[task.Id] = &task
	return &task, nil
}

func (mock *MTaskRepository) Delete(task entity.Task) (*entity.Task, error) {
	mock.tasks[task.Id] = nil
	return &task, nil
}
