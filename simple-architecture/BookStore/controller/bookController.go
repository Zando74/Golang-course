package controller

import (
	"clean-project/BookStore/interfaces"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type BookController struct {
	bookService interfaces.BookService
}

func (bc *BookController) PrintAuthor(echoCtx echo.Context) error {
	fmt.Println("** Inside Book Controller **")
	requestCtx := echoCtx.Request().Context()
	bookTitle := bc.bookService.PrintBookTitle(requestCtx)
	return echoCtx.String(http.StatusOK, bookTitle)
}

func NewBookController(echoCtx *echo.Echo, bookServiceObj interfaces.BookService) {
	bookController := &BookController{
		bookService: bookServiceObj,
	}

	echoCtx.GET("/printAuthor", bookController.PrintAuthor)
}
