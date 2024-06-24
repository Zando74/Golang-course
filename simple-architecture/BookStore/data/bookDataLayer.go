package data

import (
	"clean-project/BookStore/interfaces"
	"clean-project/BookStore/model"
	"context"
	"database/sql"
)

type BookDataLayerImpl struct {
	Dbconn *sql.DB
}

func (bdl *BookDataLayerImpl) GetBookByID(ctx context.Context, bookID int16) *model.Book {
	// make query to DB and return a book
	return &model.Book{Title: "fake book", Author: "fake Author"}
}

func NewBookDataLayerImpl(conn *sql.DB) interfaces.BookDataLayer {
	return &BookDataLayerImpl{}
}
