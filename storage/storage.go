package storage

import (
	"context"

	books "github.com/celestebrant/library-of-books/books"
)

type Storage interface {
	CreateBook(context.Context, *books.Book) (*books.Book, error)
}
