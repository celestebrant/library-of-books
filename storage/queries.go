package storage

import (
	"context"
	"fmt"
	"time"

	books "github.com/celestebrant/library-of-books/books"
	"github.com/oklog/ulid/v2"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// CreateBook creates a record in the books table (representing a book) in
// the db. If book creation time has a zero value, then it will be set to now.
// If the book ID is empty, then it will take on a generated ULID string.
func (s *MysqlStorage) CreateBook(ctx context.Context, b *books.Book) (*books.Book, error) {
	query := "INSERT INTO `books` (`id`, `creation_time`, `title`, `author`) VALUES (?, ?, ?, ?);"

	// TODO: decide if this logic should be moved elsewhere. Add unit tests.
	var creationTime time.Time
	if b.CreationTime.AsTime().IsZero() {
		creationTime = time.Now()
	} else {
		creationTime = b.CreationTime.AsTime()
	}

	if b.Id == "" {
		b.Id = ulid.Make().String()
	}

	_, err := s.db.ExecContext(ctx, query, b.Id, creationTime, b.Title, b.Author)
	if err != nil {
		return &books.Book{}, fmt.Errorf("error during insert: %w", err)
	}

	return b, nil
}

// GetBook fetches a record (book) in the books table from the db by ID.
func (s *MysqlStorage) GetBook(ctx context.Context, bookID string) (books.Book, error) {
	var id, author, title string
	var creationTimeDB []uint8
	query := "SELECT id, author, title, creation_time FROM books WHERE id = ? ;"

	row := s.db.QueryRowContext(ctx, query, bookID)
	err := row.Scan(&id, &author, &title, &creationTimeDB)
	if err != nil {
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
