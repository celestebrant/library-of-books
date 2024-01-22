package storage

import (
	"context"
	"fmt"
	"time"

	books "github.com/celestebrant/library-of-books/books"
	"github.com/oklog/ulid/v2"
)

// CreateBook creates a record in the books table (representing a book) in
// the db. If book creation time has a zero value, then it will be set to now.
// If the book ID is empty, then it will take on a generated ULID string.
func (s *MysqlStorage) CreateBook(ctx context.Context, b *books.Book) (*books.Book, error) {
	query := "INSERT INTO `books` (`id`, `creation_time`, `title`, `author`) VALUES (?, ?, ?, ?);"

	var creationTime time.Time
	if b.CreationTime.AsTime().IsZero() {
		creationTime = time.Now()
	} else {
		creationTime = b.CreationTime.AsTime()
	}

	if b.Id == "" {
		b.Id = ulid.Make().String()
	}

	// TODO: Move this validation logic elsewhere
	// err := validate_book.Validate(b)
	// if err != nil {
	// 	return nil, fmt.Errorf("cannot insert book: %w", err)
	// }

	_, err := s.db.ExecContext(ctx, query, b.Id, creationTime, b.Title, b.Author)
	if err != nil {
		return &books.Book{}, fmt.Errorf("error during insert: %w", err)
	}

	return b, nil
}
