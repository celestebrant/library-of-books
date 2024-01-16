package storage

import (
	"context"
	"fmt"
	"time"

	"github.com/celestebrant/library-of-books/book"
	"github.com/oklog/ulid/v2"
)

// CreateBook creates a record of a book in the Mysql storage. If book creation time has a zero value,
// then it will be set to now. If the book ID is empty, then it will take on a generated ULID string.
func (s *MysqlStorage) CreateBook(ctx context.Context, b *book.Book) (*book.Book, error) {
	query := "INSERT INTO `books` (`id`, `creation_time`, `title`, `author`) VALUES (?, ?, ?, ?);"

	if b.CreationTime.IsZero() {
		b.CreationTime = time.Now()
	}
	if b.Id == "" {
		b.Id = ulid.Make().String()
	}

	err := b.Validate()
	if err != nil {
		return nil, fmt.Errorf("cannot insert book: %w", err)
	}

	_, err = s.db.ExecContext(ctx, query, b.Id, b.CreationTime, b.Title, b.Author)
	if err != nil {
		return &book.Book{}, fmt.Errorf("error during insert: %w", err)
	}

	return b, nil
}
