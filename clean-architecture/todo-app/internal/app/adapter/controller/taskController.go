package controller

import (
	"net/http"
	"todo-app/internal/app/adapter/repository"
	"todo-app/internal/app/application/usecase"
	"todo-app/internal/app/domain/entity"
	"todo-app/internal/app/domain/valueobject"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

var (
	taskRepositoryImpl = repository.TaskRepository{}
)

type ErrorResponse struct {
	Message string `json:"message"`
}

type SuccessResponse struct {
	Task entity.Task `json:"task"`
}

type SuccessIndexResponse struct {
	Tasks []*entity.Task `json:"tasks"`
}

type TaskController struct{}

func TaskRouter(echoCtx *echo.Echo) {

	taskController := TaskController{}

	echoCtx.GET("/index", taskController.indexTask)
	echoCtx.POST("/create", ValidateRequestBody(taskController.createTask))
	echoCtx.POST("/update", ValidateRequestBody(taskController.updateTask))
	echoCtx.DELETE("/delete", ValidateRequestBody(taskController.deleteTask))
}

func (ctrl TaskController) indexTask(echoCtx echo.Context) error {

	tasks, err := usecase.IndexTaskUseCase(usecase.IndexTaskInputUseCase{}, taskRepositoryImpl)

	if err != nil {
		return echoCtx.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
	}

	return echoCtx.JSON(http.StatusOK, SuccessIndexResponse{Tasks: tasks})
}

func (ctrl TaskController) createTask(echoCtx echo.Context) error {

	newId := uuid.NewString()

	var newTask NewTask

	if err := echoCtx.Bind(&newTask); err != nil {
		return echoCtx.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
	}

	createdTask, err := usecase.CreateTaskUseCase(usecase.CreateTaskInputUseCase{
		Id:          newId,
		Title:       newTask.Title,
		Description: newTask.Description,
	}, taskRepositoryImpl)

	if err != nil {
		return echoCtx.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
	}

	return echoCtx.JSON(http.StatusCreated, SuccessResponse{Task: *createdTask})
}

func (ctrl TaskController) updateTask(echoCtx echo.Context) error {
	var existingTask ExistingTask

	if err := echoCtx.Bind(&existingTask); err != nil {
		return echoCtx.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
	}

	updatedTask, err := usecase.UpdateTaskUseCase(usecase.UpdateTaskInputUseCase{
		Id:          existingTask.Id,
		Title:       existingTask.Title,
		Description: existingTask.Description,
		Status:      valueobject.BaseTaskState(*existingTask.Status),
	}, taskRepositoryImpl)

	if err != nil {
		return echoCtx.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
	}

	return echoCtx.JSON(http.StatusAccepted, SuccessResponse{Task: *updatedTask})
}

func (ctrl TaskController) deleteTask(echoCtx echo.Context) error {
	var existingTask ExistingTask

	if err := echoCtx.Bind(&existingTask); err != nil {
		return echoCtx.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
	}

	deletedTask, err := usecase.DeleteTaskUseCase(usecase.DeleteTaskInputUseCase{
		Id:          existingTask.Id,
		Title:       existingTask.Title,
		Description: existingTask.Description,
		Status:      valueobject.BaseTaskState(*existingTask.Status),
	}, taskRepositoryImpl)

	if err != nil {
		return echoCtx.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
	}

	return echoCtx.JSON(http.StatusAccepted, SuccessResponse{Task: *deletedTask})
}
