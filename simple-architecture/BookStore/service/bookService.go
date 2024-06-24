package service

import (
	"clean-project/BookStore/interfaces"
	"context"
	"fmt"
)

type BookServiceImpl struct {
	BookDataLayer interfaces.BookDataLayer
}

func (bs *BookServiceImpl) PrintBookTitle(ctx context.Context) string {
	fmt.Println("** Inside Book Service **")
	// Do business logic...
	book := bs.BookDataLayer.GetBookByID(ctx, 0)
	fmt.Printf(" Book title : %s , Book Author : %s", book.Title, book.Author)
	return book.Title
}

func NewBookServiceImpl(bookDL interfaces.BookDataLayer) interfaces.BookService {
	return &BookServiceImpl{BookDataLayer: bookDL}
}
