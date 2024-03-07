package storage

import (
	"context"
	"fmt"
	"time"

	books "github.com/celestebrant/library-of-books/books"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// CreateBook creates a record in the books table (representing a book) in
// the db.
func (s *MysqlStorage) CreateBook(ctx context.Context, b *Book) error {
	query := "INSERT INTO `books` (`id`, `creation_time`, `title`, `author`) VALUES (?, ?, ?, ?);"

	if _, err := s.db.ExecContext(ctx, query, b.Id, b.CreationTime, b.Title, b.Author); err != nil {
		return fmt.Errorf("error during insert: %w", err)
	}

	return nil
}

// GetBook fetches a record (book) in the books table from the db by ID.
func (s *MysqlStorage) GetBook(ctx context.Context, bookID string) (books.Book, error) {
	var id, author, title string
	var creationTimeDB []uint8
	query := "SELECT id, author, title, creation_time FROM books WHERE id = ? ;"

	row := s.db.QueryRowContext(ctx, query, bookID)
	if err := row.Scan(&id, &author, &title, &creationTimeDB); err != nil {
		return books.Book{}, err
	}

	creationTime, err := time.Parse(time.DateTime, string(creationTimeDB))
	if err != nil {
		return books.Book{}, fmt.Errorf("cannot parse creation_time: %w", err)
	}

	return books.Book{
		Id:           id,
		Author:       author,
		Title:        title,
		CreationTime: timestamppb.New(creationTime),
	}, nil
}
