package controller

/* INTEGRATION TEST */

import (
	"clean-project/BookStore/data"
	"clean-project/BookStore/service"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestPrintAuthor(t *testing.T) {
	// setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/PrintAuthor", strings.NewReader(""))
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	dataLayer := data.NewBookDataLayerImpl(nil)
	service := service.NewBookServiceImpl(dataLayer)
	handler := &BookController{service}

	// assertion
	if assert.NoError(t, handler.PrintAuthor(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "fake book", rec.Body.String())
	}
}
