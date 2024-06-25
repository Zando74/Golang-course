package domain_test

import (
	"testing"
	"todo-app/internal/app/domain/valueobject"
)

func UnexpectedErrorAssert(err error, t *testing.T) {
	if err != nil {
		t.Errorf("Unexpected Error occured : %s", err)
	}
}
func ExpectedErrorAssert(currentErr error, expectedError error, t *testing.T) {
	if currentErr == nil {
		t.Errorf("Error was expected : %s", expectedError)
	}
}

func DifferentIdAssert(createdTaskId, taskId string, t *testing.T) {
	if createdTaskId != taskId {
		t.Errorf("Unexpected Task Id")
	}
}

func DifferentTitleAssert(createdTaskTitle, taskTitle string, t *testing.T) {
	if createdTaskTitle != taskTitle {
		t.Errorf("Unexpected Task Title")
	}
}

func DifferentDescriptionAssert(createdTaskDescription, taskDescription string, t *testing.T) {
	if createdTaskDescription != taskDescription {
		t.Errorf("Unexpected Task Description")
	}
}

func UnexpectedStateAssert(initialState, newState valueobject.TaskState, t *testing.T) {
	if initialState.String() != newState.String() {
		t.Error("Unexpected State")
	}
}
