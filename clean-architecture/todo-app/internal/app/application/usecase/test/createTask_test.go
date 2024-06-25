package usecase_test

import (
	"testing"
	"todo-app/internal/app/application/usecase"
	domain_test "todo-app/internal/app/domain/test"
	"todo-app/internal/app/domain/valueobject"
	"todo-app/internal/app/test/mock"
)

func TestValidTaskCreation(t *testing.T) {
	mTaskRepo := mock.NewMTaskRepository()

	taskId := "task-id"
	taskTitle := "A task title"
	taskDescription := "A tast Description"

	createdTask, err := usecase.CreateTaskUseCase(usecase.CreateTaskInputUseCase{Id: taskId, Title: taskTitle, Description: taskDescription}, mTaskRepo)

	domain_test.UnexpectedErrorAssert(err, t)
	domain_test.DifferentIdAssert(createdTask.Id, taskId, t)
	domain_test.DifferentTitleAssert(createdTask.Title.String(), taskTitle, t)
	domain_test.DifferentDescriptionAssert(createdTask.Description.String(), taskDescription, t)

}

func TestTooLongTitleTaskCreation(t *testing.T) {
	mTaskRepo := mock.NewMTaskRepository()

	taskId := "task-id"
	taskTitle := "A"

	for len(taskTitle) < valueobject.TITLE_MAX_LENGTH {
		taskTitle += "A"
	}
	taskDescription := "A tast Description"

	_, err := usecase.CreateTaskUseCase(usecase.CreateTaskInputUseCase{Id: taskId, Title: taskTitle, Description: taskDescription}, mTaskRepo)

	domain_test.ExpectedErrorAssert(err, &valueobject.TooLongTaskTitleError{}, t)

}

func TestTooLongDescriptionTaskCreation(t *testing.T) {
	mTaskRepo := mock.NewMTaskRepository()

	taskId := "task-id"
	taskTitle := "A task title"
	taskDescription := "A"

	for len(taskDescription) < valueobject.DESCRIPTION_MAX_LENGTH {
		taskDescription += "A"
	}

	_, err := usecase.CreateTaskUseCase(usecase.CreateTaskInputUseCase{Id: taskId, Title: taskTitle, Description: taskDescription}, mTaskRepo)

	domain_test.ExpectedErrorAssert(err, &valueobject.TooLongTaskDescriptionError{}, t)
}
