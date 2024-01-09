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

type InvalidAuthorError struct {
	Message string
}

func (e *InvalidAuthorError) Error() string {
	return e.Message
}

type InvalidTitleError struct {
	Message string
}

func (e *InvalidTitleError) Error() string {
	return e.Message
}

type InvalidIDError struct {
	Message string
}

func (e *InvalidIDError) Error() string {
	return e.Message
}

type InvalidCreationTimeError struct {
	Message string
}

func (e *InvalidCreationTimeError) Error() string {
	return e.Message
}

func (b *Book) validateAuthor() error {
	if len(b.Author) > 255 {
		return &InvalidAuthorError{"author exceeds 255 characters"}
	}
	if len(b.Author) == 0 {
		return &InvalidAuthorError{"author is empty"}
	}
	return nil
}

func (b *Book) validateTitle() error {
	if len(b.Title) > 255 {
		return &InvalidTitleError{"title exceeds 255 characters"}
	}
	if len(b.Title) == 0 {
		return &InvalidTitleError{"title is empty"}
	}
	return nil
}

func (b *Book) validateID() error {
	if len(b.ID) > 26 {
		return &InvalidIDError{"id exceeds 255 characters"}
	}
	if len(b.ID) == 0 {
		return &InvalidIDError{"id is empty"}
	}
	return nil
}

func (b *Book) validateCreationTime() error {
	if b.CreationTime.IsZero() {
		return &InvalidCreationTimeError{"creation time has zero value"}
	}
	return nil
}

// Validate validates ID, author and title of a Book. Returns an error if fields
// author, title, or ID are empty or exceed their max length. Returns an error if
// creation time has zero-value.
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
