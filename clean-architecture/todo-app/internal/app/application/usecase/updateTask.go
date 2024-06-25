package usecase

import (
	"todo-app/internal/app/domain/entity"
	"todo-app/internal/app/domain/factory"
	"todo-app/internal/app/domain/repository"
	"todo-app/internal/app/domain/valueobject"
)

type UpdateTaskInputUseCase struct {
	Id, Title, Description string
	Status                 valueobject.BaseTaskState
}

func UpdateTaskUseCase(input UpdateTaskInputUseCase, repository repository.ITask) (*entity.Task, error) {

	existingTask, err := repository.Get(input.Id)

	if err != nil {
		return nil, err
	}

	_, err = existingTask.Status.GoTo(input.Status)

	if err != nil {
		return nil, err
	}

	newTask, err := factory.TaskFactory{}.GenerateAtGivenState(input.Id, input.Title, input.Description, input.Status)
	if err != nil {
		return nil, err
	}

	updatedTask, err := repository.Update(*newTask)
	if err != nil {
		return nil, err
	}

	return updatedTask, nil
}
