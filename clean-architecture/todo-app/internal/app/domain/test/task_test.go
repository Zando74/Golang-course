package domain_test

import (
	"testing"
	"todo-app/internal/app/domain/factory"
	"todo-app/internal/app/domain/valueobject"
)

/* UNIT TEST */

func TestTaskGeneration(t *testing.T) {
	task, err := factory.TaskFactory{}.Generate("fake-id", "fake title", "fake description")

	UnexpectedErrorAssert(err, t)
	DifferentIdAssert(task.Id, "fake-id", t)
	DifferentTitleAssert(task.Title.String(), "fake title", t)
	DifferentDescriptionAssert(task.Description.String(), "fake description", t)
	UnexpectedStateAssert(task.Status, valueobject.TaskState(valueobject.Todo), t)
}
