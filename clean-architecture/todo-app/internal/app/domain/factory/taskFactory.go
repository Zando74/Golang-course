package factory

import (
	"todo-app/internal/app/domain/entity"
	"todo-app/internal/app/domain/valueobject"
)

type TaskFactory struct{}

func (tf TaskFactory) Generate(id, title, description string) (*entity.Task, error) {

	validTitle, err := valueobject.NewTaskTitle(title)
	if err != nil {
		return nil, err
	}

	validDescription, err := valueobject.NewTaskDescription(description)
	if err != nil {
		return nil, err
	}

	initialState, err := valueobject.NewState(valueobject.Todo)
	if err != nil {
		return nil, err
	}

	return &entity.Task{
			Id:          id,
			Title:       *validTitle,
			Description: *validDescription,
			Status:      *initialState},
		nil

}

func (tf TaskFactory) GenerateAtGivenState(id, title, description string, state valueobject.BaseTaskState) (*entity.Task, error) {
	validTitle, err := valueobject.NewTaskTitle(title)
	if err != nil {
		return nil, err
	}

	validDescription, err := valueobject.NewTaskDescription(description)
	if err != nil {
		return nil, err
	}

	givenState, err := valueobject.NewState(state)
	if err != nil {
		return nil, err
	}

	return &entity.Task{Id: id,
			Title:       *validTitle,
			Description: *validDescription,
			Status:      *givenState},
		nil
}
