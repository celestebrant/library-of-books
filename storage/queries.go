package storage

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/celestebrant/library-of-books/books"
	"github.com/celestebrant/library-of-books/utils"
)

// CreateBook inserts a new book record into the 'books' table using the provided Book struct.
// It takes a context for cancellation and a pointer to a Book struct containing the new book's details.
// Returns an error if the insert operation fails, including context about the failure.
func (s *MysqlStorage) CreateBook(ctx context.Context, b *Book) error {
	query := "INSERT INTO `books` (`id`, `creation_time`, `title`, `author`) VALUES (?, ?, ?, ?);"

	if _, err := s.db.ExecContext(ctx, query, b.Id, b.CreationTime, b.Title, b.Author); err != nil {
		return fmt.Errorf("failed perform SQL query: %w", err)
	}

	return nil
}

func (s *MysqlStorage) ListBooks(
	ctx context.Context, author, title string, pageSize int64, pageToken string,
) (*books.ListBooksResponse, error) {
	// No need to order because
	query := `SELECT id, title, author, creation_time
	FROM books
	WHERE (author LIKE CONCAT('%', ?, '%') OR ? IS NULL)
	  AND (title LIKE CONCAT('%', ?, '%') OR ? IS NULL)
	ORDER BY creation_time ASC
	LIMIT ?  -- page size
	OFFSET ?; -- skip this number of preceding rows
	`
	offset, err := utils.Offset(pageToken)
	if err != nil {
		return nil, err
	}

	rows, err := s.db.Query(query, author, author, title, title, pageSize, offset)
	if err != nil {
		return nil, fmt.Errorf("failed perform SQL query: %w", err)
	}
	defer rows.Close()

	// Convert the rows into list of books
	var fetchedBooks []*books.Book
	for rows.Next() {
		var b books.Book
		var creationTimeDB []uint8
		if err := rows.Scan(&b.Id, &b.Title, &b.Author, &creationTimeDB); err != nil {
			return nil, fmt.Errorf("failed to parse row into Book: %w", err)
		}

		creationTime, err := time.Parse(time.DateTime, string(creationTimeDB))
		if err != nil {
			return nil, fmt.Errorf("failed to parse creation time from []uint8 to time.Time from ListBooks SQL query: %w", err)
		}
		b.CreationTime = timestamppb.New(creationTime)

		fetchedBooks = append(fetchedBooks, &b)
	}

	// Iteration errors
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error encountered when iterating over rows: %w", err)
	}

	// Generate next page token if more results exist
	var nextPageToken string
	if len(fetchedBooks) == int(pageSize) {
		nextPageToken, err = utils.NextPageToken(pageToken, pageSize)
		if err != nil {
			return nil, err
		}
	}

	return &books.ListBooksResponse{
		Books:         fetchedBooks,
		NextPageToken: nextPageToken,
	}, nil
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
