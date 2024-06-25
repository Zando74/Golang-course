package usecase_test

import (
	"testing"
	"todo-app/internal/app/application/usecase"
	domain_test "todo-app/internal/app/domain/test"
	"todo-app/internal/app/domain/valueobject"
	"todo-app/internal/app/test/mock"
)

func TestValidTaskUpdate(t *testing.T) {
	mTaskRepo := mock.NewMTaskRepository()
	taskId := "task-id"
	taskTitle := "A task title"
	taskDescription := "A tast Description"
	createdTask, _ := usecase.CreateTaskUseCase(usecase.CreateTaskInputUseCase{Id: taskId, Title: taskTitle, Description: taskDescription}, mTaskRepo)

	newTitle, _ := valueobject.NewTaskTitle("a different title")
	newDescription, _ := valueobject.NewTaskTitle("a different Description")
	newState, _ := valueobject.NewState(valueobject.InProgress)

	updatedTask, err := usecase.UpdateTaskUseCase(usecase.UpdateTaskInputUseCase{Id: createdTask.Id, Title: newTitle.String(), Description: newDescription.String(), Status: newState.Get()}, mTaskRepo)

	domain_test.UnexpectedErrorAssert(err, t)

	domain_test.DifferentIdAssert(createdTask.Id, taskId, t)
	domain_test.DifferentTitleAssert(createdTask.Title.String(), taskTitle, t)
	domain_test.DifferentDescriptionAssert(createdTask.Description.String(), taskDescription, t)

	domain_test.UnexpectedStateAssert(updatedTask.Status, *newState, t)
}

func TestInvalidStateUpdate(t *testing.T) {
	mTaskRepo := mock.NewMTaskRepository()
	taskId := "task-id"
	taskTitle := "A task title"
	taskDescription := "A tast Description"
	createdTask, _ := usecase.CreateTaskUseCase(usecase.CreateTaskInputUseCase{Id: taskId, Title: taskTitle, Description: taskDescription}, mTaskRepo)

	newTitle, _ := valueobject.NewTaskTitle("a different title")
	newDescription, _ := valueobject.NewTaskTitle("a different Description")
	newState, _ := valueobject.NewState(valueobject.Done)

	_, err := usecase.UpdateTaskUseCase(usecase.UpdateTaskInputUseCase{Id: createdTask.Id, Title: newTitle.String(), Description: newDescription.String(), Status: newState.Get()}, mTaskRepo)

	domain_test.ExpectedErrorAssert(err, &valueobject.InvalidTaskStateTransitionError{Current: &createdTask.Status, Desired: newState}, t)

}
