package book

import (
	"time"
)

type Book struct {
	ID           string
	Title        string
	Author       string
	CreationTime time.Time
}

// validateAuthor returns an error if author exceeds 255 characters or is empty.
func (b *Book) validateAuthor() error {
	if len(b.Author) > 255 {
		return &InvalidAuthorError{"author exceeds 255 characters"}
	}
	if len(b.Author) == 0 {
		return &InvalidAuthorError{"author is empty"}
	}
	return nil
}

// validateTitle returns an error if title exceeds 255 characters or is empty.
func (b *Book) validateTitle() error {
	if len(b.Title) > 255 {
		return &InvalidTitleError{"title exceeds 255 characters"}
	}
	if len(b.Title) == 0 {
		return &InvalidTitleError{"title is empty"}
	}
	return nil
}

// validateID returns an error if ID exceeds 255 characters or is empty.
func (b *Book) validateID() error {
	if len(b.ID) > 26 {
		return &InvalidIDError{"id exceeds 255 characters"}
	}
	if len(b.ID) == 0 {
		return &InvalidIDError{"id is empty"}
	}
	return nil
}

// validateCreationTime returns an error if creation time has a zero-value.
func (b *Book) validateCreationTime() error {
	if b.CreationTime.IsZero() {
		return &InvalidCreationTimeError{"creation time has zero value"}
	}
	return nil
}

// Validate validates ID, author, title, and creation timestamp of a Book.
func (b *Book) Validate() error {
	err := b.validateAuthor()
	if err != nil {
		return err
	}

	err = b.validateTitle()
	if err != nil {
		return err
	}

	err = b.validateID()
	if err != nil {
		return err
	}

	err = b.validateCreationTime()
	if err != nil {
		return err
	}

	return nil
}
