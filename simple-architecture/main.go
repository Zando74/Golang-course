package main

/*
	Bookstore application
*/

import (
	"clean-project/BookStore/controller"
	"clean-project/BookStore/data"
	"clean-project/BookStore/service"

	"github.com/labstack/echo/v4"
)

func main() {
	echoContext := echo.New()

	dataLayer := data.NewBookDataLayerImpl(nil)
	service := service.NewBookServiceImpl(dataLayer)
	controller.NewBookController(echoContext, service)

	echoContext.Logger.Info(echoContext.Start(":8080"))
}
