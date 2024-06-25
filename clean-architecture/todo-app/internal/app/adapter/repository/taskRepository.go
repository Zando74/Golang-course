package repository

import (
	"fmt"
	"todo-app/internal/app/adapter/postgresql"
	"todo-app/internal/app/adapter/postgresql/model"
	"todo-app/internal/app/domain/entity"
	"todo-app/internal/app/domain/factory"
	"todo-app/internal/app/domain/valueobject"
)

type TaskRepository struct{}

func (t TaskRepository) Get(id string) (*entity.Task, error) {
	db := postgresql.Connection()
	var task model.Task
	result := db.First(&task, "id = ?", id)

	if result.Error != nil {
		return nil, fmt.Errorf("failed to retrieve task")
	}

	return factory.TaskFactory{}.GenerateAtGivenState(
		task.Id,
		task.Title,
		task.Description,
		valueobject.BaseTaskState(task.Status),
	)

}

func (t TaskRepository) Index() ([]*entity.Task, error) {
	db := postgresql.Connection()
	var tasks []model.Task

	result := db.Find(&tasks)

	if result.Error != nil {
		return nil, fmt.Errorf("failed to retrieve tasks")
	}

	taskFactory := factory.TaskFactory{}
	tasksEntities := make([]*entity.Task, 0, len(tasks))

	for _, task := range tasks {

		taskEntityCreatedFromModel, err := taskFactory.GenerateAtGivenState(
			task.Id,
			task.Title,
			task.Description,
			valueobject.BaseTaskState(task.Status),
		)

		if err != nil {
			return nil, err
		}

		tasksEntities = append(tasksEntities, taskEntityCreatedFromModel)

	}

	return tasksEntities, nil

}

func (t TaskRepository) Create(task entity.Task) (*entity.Task, error) {
	db := postgresql.Connection()
	newTask := model.NewTaskModelFromEntity(task)

	result := db.Create(&newTask)

	if result.Error != nil {
		return nil, fmt.Errorf("failed to create task")
	}

	return &task, nil
}

func (t TaskRepository) Update(task entity.Task) (*entity.Task, error) {
	db := postgresql.Connection()
	taskToUpdate := model.NewTaskModelFromEntity(task)

	result := db.Save(&taskToUpdate)

	if result.Error != nil {
		return nil, fmt.Errorf("failed to update task")
	}

	return &task, nil
}

func (t TaskRepository) Delete(task entity.Task) (*entity.Task, error) {
	db := postgresql.Connection()
	taskToDelete := model.NewTaskModelFromEntity(task)

	result := db.Delete(&taskToDelete)

	if result.Error != nil {
		return nil, fmt.Errorf("failed to delete task")
	}

	return &task, nil
}
