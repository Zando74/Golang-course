package controller

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type NewTask struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
}

type ExistingTask struct {
	Id          string `json:"id" validate:"required,uuid"`
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	Status      *int   `json:"status" validate:"required"`
}

func ValidateRequestBody(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		v := validator.New()

		bodyBytes, err := io.ReadAll(c.Request().Body)
		if err != nil {
			return err
		}
		c.Request().Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

		var body interface{}

		switch c.Request().URL.Path {
		case "/create":
			body = new(NewTask)
		case "/update":
			body = new(ExistingTask)
		case "/delete":
			body = new(ExistingTask)
		default:
			return next(c)
		}

		if err := json.Unmarshal(bodyBytes, &body); err != nil {
			return err
		}

		if err := v.Struct(body); err != nil {
			return c.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
		}

		return next(c)
	}
}
