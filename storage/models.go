package storage

import (
	"time"

	books "github.com/celestebrant/library-of-books/books"
	"github.com/oklog/ulid/v2"
)

// Book is a representation of a books.Book that has been transformed to be inserted
// into a SQL database.
type Book struct {
	Id           string
	Title        string
	Author       string
	CreationTime time.Time
}

// NewBook returns a Book with fields populated based on req. Sets CreationTime
// to time.Now (UTC) if req.Book.CreationTime is not set (default being
// 1970-01-01 00:00:00.0 UTC). Sets Id to a new ULID if req.Book.Id is an empty string.
func NewBook(req *books.CreateBookRequest) *Book {
	var creationTime time.Time
	if req.Book.CreationTime.AsTime() == time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC) {
		creationTime = time.Now()
	} else {
		creationTime = req.Book.CreationTime.AsTime()
	}

	var id string
	if req.Book.Id == "" {
		id = ulid.Make().String()
	} else {
		id = req.Book.Id
	}

	return &Book{
		Id:           id,
		Title:        req.Book.Title,
		Author:       req.Book.Author,
		CreationTime: creationTime.UTC(),
	}
}
