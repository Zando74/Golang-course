package interfaces

import (
	"clean-project/BookStore/model"
	"context"
)

type BookService interface {
	PrintBookTitle(ctx context.Context) string
}

type BookDataLayer interface {
	GetBookByID(ctx context.Context, bookID int16) *model.Book
}
