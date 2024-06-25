package usecase

import (
	"todo-app/internal/app/domain/entity"
	"todo-app/internal/app/domain/factory"
	"todo-app/internal/app/domain/repository"
	"todo-app/internal/app/domain/valueobject"
)

type DeleteTaskInputUseCase struct {
	Id, Title, Description string
	Status                 valueobject.BaseTaskState
}

func DeleteTaskUseCase(input DeleteTaskInputUseCase, repository repository.ITask) (*entity.Task, error) {
	taskToDelete, err := factory.TaskFactory{}.GenerateAtGivenState(input.Id, input.Title, input.Description, input.Status)
	if err != nil {
		return nil, err
	}

	deletedTask, err := repository.Delete(*taskToDelete)
	if err != nil {
		return nil, err
	}

	return deletedTask, nil
}
