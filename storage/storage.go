package storage

import (
	"context"

	"github.com/celestebrant/library-of-books/book"
)

type Storage interface {
	CreateBook(context.Context, *book.Book) (*book.Book, error)
}
