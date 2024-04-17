package booksservice

import (
	"fmt"
	"time"

	books "github.com/celestebrant/library-of-books/books"
)

type Book struct {
	Id           string
	Title        string
	Author       string
	CreationTime time.Time
}

const (
	idMaxLength        = 30
	requestIDMaxLength = 30
	authorMaxLength    = 255
	titleMaxLength     = 255
	pageSizeMaxLength  = 50
)

/*
ValidateCreateBookRequest returns an error if the following are not satisfied:
- Request ID is not empty and does not exceed the maximum allowed length.
- Author field within the Book struct is not empty and does not exceed the maximum allowed length.
- Title field within the Book struct is not empty and does not exceed the maximum allowed length.
- Id field within the Book struct (optional) does not exceed the maximum allowed length (if a length limit exists for ID).
*/
func ValidateCreateBookRequest(req *books.CreateBookRequest) error {
	if len(req.RequestId) == 0 {
		return &ValidationError{
			Field:   "request_id",
			Message: "must not be empty",
		}
	} else if len(req.RequestId) > requestIDMaxLength {
		return &ValidationError{
			Field:   "request_id",
			Message: fmt.Sprintf("must not exceed %d characters", requestIDMaxLength),
		}
	}

	if len(req.Book.Author) == 0 {
		return &ValidationError{
			Field:   "author",
			Message: "must not be empty",
		}
	} else if len(req.Book.Author) > authorMaxLength {
		return &ValidationError{
			Field:   "author",
			Message: fmt.Sprintf("must not exceed %d characters", authorMaxLength),
		}
	}

	if len(req.Book.Title) == 0 {
		return &ValidationError{
			Field:   "title",
			Message: "must not be empty",
		}
	} else if len(req.Book.Title) > titleMaxLength {
		return &ValidationError{
			Field:   "title",
			Message: fmt.Sprintf("must not exceed %d characters", titleMaxLength),
		}
	}

	if len(req.Book.Id) > idMaxLength {
		return &ValidationError{
			Field:   "id",
			Message: fmt.Sprintf("must not exceed %d characters", idMaxLength),
		}
	}

	return nil
}

// ValidateListBooksRequest returns an error if page size is outside limits (1 - 50).
func ValidateListBooksRequest(req *books.ListBooksRequest) error {
	if req.PageSize <= 0 || req.PageSize > 50 {
		return &ValidationError{
			Field:   "page_size",
			Message: fmt.Sprintf("must be greater than zero and not exceed %d", pageSizeMaxLength),
		}
	}
	return nil
}
