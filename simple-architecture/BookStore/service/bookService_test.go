package service

/* UNIT TEST */

import (
	"clean-project/BookStore/model"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

type BookDataLayerTestImpl struct {
}

func (bdlTest *BookDataLayerTestImpl) GetBookByID(ctx context.Context, bookID int16) *model.Book {
	return &model.Book{Title: "fake title", Author: "fake Author"}
}

func TestPrintBookTitle(t *testing.T) {
	bookDataLayer := &BookDataLayerTestImpl{}
	bookServiceImpl := NewBookServiceImpl(bookDataLayer)
	bookTitle := bookServiceImpl.PrintBookTitle(context.Background())

	expectedTitle := "fake title"

	assert.Equal(t, expectedTitle, bookTitle)

}
