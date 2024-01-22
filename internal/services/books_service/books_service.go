package books_service

import (
	"context"
	"fmt"

	books "github.com/celestebrant/library-of-books/books"
)

type BooksServer struct {
	books.UnimplementedBooksServer
}

func (s *BooksServer) CreateBook(
	ctx context.Context, req *books.CreateBookRequest,
) (*books.CreateBookResponse, error) {
	err := Validate(req.Book)
	if err != nil {
		return nil, fmt.Errorf("validation error: %w", err)
	}

	res := &books.CreateBookResponse{Book: req.Book}
	return res, nil
}
