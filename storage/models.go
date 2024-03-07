package storage

import (
	"time"

	books "github.com/celestebrant/library-of-books/books"
	"github.com/oklog/ulid/v2"
)

// Book defines the schema for a book record suitable for storage in an SQL database,
// including a unique identifier, title, author, and creation time.
type Book struct {
	Id           string
	Title        string
	Author       string
	CreationTime time.Time
}

// NewBookFromRequest constructs a Book instance from a CreateBookRequest. It assigns a 
// current UTC timestamp to CreationTime if unspecified, and generates a new ULID for Id
// if empty. The Title and Author fields are directly mapped from the request.
func NewBookFromRequest(req *books.CreateBookRequest) *Book {
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
