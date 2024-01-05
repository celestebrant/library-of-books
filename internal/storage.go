package storage

import "github.com/celestebrant/docker-sql-demo/book"

type Storage interface {
	Create(book.Book) error
}
