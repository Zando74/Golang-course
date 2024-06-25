package usecase

import (
	"todo-app/internal/app/domain/entity"
	"todo-app/internal/app/domain/factory"
	"todo-app/internal/app/domain/repository"
)

type CreateTaskInputUseCase struct {
	Id, Title, Description string
}

func CreateTaskUseCase(input CreateTaskInputUseCase, repository repository.ITask) (*entity.Task, error) {
	newTask, err := factory.TaskFactory{}.Generate(input.Id, input.Title, input.Description)
	if err != nil {
		return nil, err
	}

	createdTask, err := repository.Create(*newTask)
	if err != nil {
		return nil, err
	}

	return createdTask, nil
}
