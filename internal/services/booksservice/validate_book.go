package booksservice

import (
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
	idMaxLength     = 30
	authorMaxLength = 255
	titleMaxLength  = 255
)

// validateAuthor returns an error if author exceeds max characters or is empty.
func validateAuthor(b *books.Book) error {
	if len(b.Author) > authorMaxLength {
		return &InvalidAuthorError{"author cannot exceed 255 characters"}
	}
	if len(b.Author) == 0 {
		return &InvalidAuthorError{"author cannot be empty"}
	}
	return nil
}

// validateTitle returns an error if title exceeds max characters or is empty.
func validateTitle(b *books.Book) error {
	if len(b.Title) > titleMaxLength {
		return &InvalidTitleError{"title cannot exceed 255 characters"}
	}
	if len(b.Title) == 0 {
		return &InvalidTitleError{"title cannot be empty"}
	}
	return nil
}

// validateID returns an error if ID exceeds max characters.
func validateID(b *books.Book) error {
	if len(b.Id) > idMaxLength {
		return &InvalidIDError{"id cannot exceed 30 characters"}
	}
	return nil
}

// Validate returns an error for semantically invalid fields:
// Id, Author, Title.
func Validate(b *books.Book) error {
	if err := validateAuthor(b); err != nil {
		return err
	}

	if err := validateTitle(b); err != nil {
		return err
	}

	if err := validateID(b); err != nil {
		return err
	}

	return nil
}

// ValidateListBooksRequest returns an error if page size is not within range 1 - 50, inclusive.
func ValidateListBooksRequest(req *books.ListBooksRequest) error {
	if req.PageSize <= 0 || req.PageSize > 50 {
		return &InvalidPageSizeError{"page size must be greater than 0 and less than or equal to 50"}
	}
	return nil
}
