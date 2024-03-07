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

// validateAuthor returns an error if author exceeds 255 characters or is empty.
func validateAuthor(b *books.Book) error {
	if len(b.Author) > 255 {
		return &InvalidAuthorError{"author cannot exceed 255 characters"}
	}
	if len(b.Author) == 0 {
		return &InvalidAuthorError{"author cannot be empty"}
	}
	return nil
}

// validateTitle returns an error if title exceeds 255 characters or is empty.
func validateTitle(b *books.Book) error {
	if len(b.Title) > 255 {
		return &InvalidTitleError{"title cannot exceed 255 characters"}
	}
	if len(b.Title) == 0 {
		return &InvalidTitleError{"title cannot be empty"}
	}
	return nil
}

// validateID returns an error if ID exceeds 26 characters or is empty.
func validateID(b *books.Book) error {
	if len(b.Id) > 26 {
		return &InvalidIDError{"id cannot exceed 26 characters"}
	}
	if len(b.Id) == 0 {
		return &InvalidIDError{"id cannot be empty"}
	}
	return nil
}

// Validate returns an error for semantically invalid fields for b:
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
