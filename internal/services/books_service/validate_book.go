package books_service

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
		return &InvalidAuthorError{"author exceeds 255 characters"}
	}
	if len(b.Author) == 0 {
		return &InvalidAuthorError{"author is empty"}
	}
	return nil
}

// validateTitle returns an error if title exceeds 255 characters or is empty.
func validateTitle(b *books.Book) error {
	if len(b.Title) > 255 {
		return &InvalidTitleError{"title exceeds 255 characters"}
	}
	if len(b.Title) == 0 {
		return &InvalidTitleError{"title is empty"}
	}
	return nil
}

// validateID returns an error if ID exceeds 255 characters or is empty.
func validateID(b *books.Book) error {
	if len(b.Id) > 26 {
		return &InvalidIDError{"id exceeds 255 characters"}
	}
	if len(b.Id) == 0 {
		return &InvalidIDError{"id is empty"}
	}
	return nil
}

// validateCreationTime returns an error if creation time has a zero-value.
func validateCreationTime(b *books.Book) error {
	time := b.CreationTime.AsTime()
	if time.IsZero() {
		return &InvalidCreationTimeError{"creation time has zero value"}
	}
	return nil
}

// Validate validates ID, author, title, and creation timestamp of a Book.
func Validate(b *books.Book) error {
	err := validateAuthor(b)
	if err != nil {
		return err
	}

	err = validateTitle(b)
	if err != nil {
		return err
	}

	err = validateID(b)
	if err != nil {
		return err
	}

	err = validateCreationTime(b)
	if err != nil {
		return err
	}

	return nil
}
