package storage

import (
	"context"
	"fmt"
	"time"
)

// CreateBook inserts a new book record into the 'books' table using the provided Book struct.
// It takes a context for cancellation and a pointer to a Book struct containing the new book's details.
// Returns an error if the insert operation fails, including context about the failure.
func (s *MysqlStorage) CreateBook(ctx context.Context, b *Book) error {
	query := "INSERT INTO `books` (`id`, `creation_time`, `title`, `author`) VALUES (?, ?, ?, ?);"

	if _, err := s.db.ExecContext(ctx, query, b.Id, b.CreationTime, b.Title, b.Author); err != nil {
		return fmt.Errorf("error during insert: %w", err)
	}

	return nil
}

// GetBook retrieves a book from the 'books' table for a given bookID. It returns a
// populated Book struct on success. It returns sql.ErrNoRows if the book is not found,
// or another error for any issues during query execution or data parsing.
func (s *MysqlStorage) GetBook(ctx context.Context, bookID string) (Book, error) {
	var id, author, title string
	var creationTimeDB []uint8
	query := "SELECT id, author, title, creation_time FROM books WHERE id = ? ;"

	row := s.db.QueryRowContext(ctx, query, bookID)
	if err := row.Scan(&id, &author, &title, &creationTimeDB); err != nil {
		return Book{}, err
	}

	creationTime, err := time.Parse(time.DateTime, string(creationTimeDB))
	if err != nil {
		return Book{}, fmt.Errorf("cannot parse creation_time: %w", err)
	}

	return Book{
		Id:           id,
		Author:       author,
		Title:        title,
		CreationTime: creationTime,
	}, nil
}
