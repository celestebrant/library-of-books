package books_service

import (
	"context"
	"fmt"

	books "github.com/celestebrant/library-of-books/books"
	validate_book "github.com/celestebrant/library-of-books/validate_book"
)

type BooksServer struct {
	books.UnimplementedBooksServer
}

func (s *BooksServer) CreateBook(
	ctx context.Context, req *books.CreateBookRequest,
) (*books.CreateBookResponse, error) {
	err := validate_book.Validate(req.Book)
	if err != nil {
		return nil, fmt.Errorf("validation error: %w", err)
	}

	res := &books.CreateBookResponse{Book: req.Book}
	return res, nil
}
