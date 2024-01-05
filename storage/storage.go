package storage

import (
	"context"

	"github.com/celestebrant/docker-sql-demo/book"
)

type Storage interface {
	CreateBook(context.Context, *book.Book) (*book.Book, error)
}
