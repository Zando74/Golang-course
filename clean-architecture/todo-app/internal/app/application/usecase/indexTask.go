package usecase

import (
	"todo-app/internal/app/domain/entity"
	"todo-app/internal/app/domain/repository"
)

type IndexTaskInputUseCase struct{}

func IndexTaskUseCase(input IndexTaskInputUseCase, repository repository.ITask) ([]*entity.Task, error) {

	tasks, err := repository.Index()
	if err != nil {
		return nil, err
	}

	return tasks, nil
}
