package books_service

import (
	"context"

	books "github.com/celestebrant/library-of-books/books"
	"github.com/celestebrant/library-of-books/storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type BooksServer struct {
	books.UnimplementedBooksServer
	*storage.MysqlStorage
}

func (s *BooksServer) CreateBook(
	ctx context.Context, req *books.CreateBookRequest,
) (*books.CreateBookResponse, error) {
	err := Validate(req.Book)

	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	// TODO: decide if returned Book should be used in res. If no, add tests to compare res.Book with db-returned book
	_, err = s.MysqlStorage.CreateBook(ctx, req.Book)

	res := &books.CreateBookResponse{Book: req.Book}
	return res, nil
}
