package tests

import (
	"context"
	"log"
	"sync"
	"testing"
	"time"

	books "github.com/celestebrant/library-of-books/books"
	"github.com/celestebrant/library-of-books/internal/services/booksclient"
	"github.com/celestebrant/library-of-books/internal/services/booksservice"
	"github.com/celestebrant/library-of-books/storage"
	"github.com/oklog/ulid/v2"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// setUpServerAndClient sets up the server and client. Returns the client and
// tear-down actvities which should be deferred.
func setUpServerAndClient(address string) (books.BooksClient, func()) {
	var wg sync.WaitGroup
	wg.Add(1)
	server, lis := booksservice.MustNewBooksServer(address, &wg)
	client, conn := booksclient.MustNewBooksClient(address)

	return client, func() {
		booksservice.StopBooksServer(server, lis, &wg)
		conn.Close()
	}
}

// TestCreateBook contains integration tests for the CreateBook endpoint.
func TestCreateBook(t *testing.T) {
	// Prepare set up and tear down of server and client.
	client, tearDown := setUpServerAndClient("127.0.0.1:8089")
	defer tearDown()

	t.Run("writes to db", func(t *testing.T) {
		r := require.New(t)

		res, err := client.CreateBook(
			context.Background(),
			&books.CreateBookRequest{
				Book: &books.Book{
					Id:           ulid.Make().String(),
					Title:        ulid.Make().String(),
					Author:       ulid.Make().String(),
					CreationTime: timestamppb.New(time.Now()),
				},
				RequestId: ulid.Make().String(),
			},
		)
		r.NoError(err)
		r.NotEmpty(res)

		// Verify against db record
		dbConnection, err := storage.NewMysqlStorage(storage.MysqlConfig{
			Username: "user1",
			Password: "password1",
			DBName:   "library",
			Port:     3306,
			Host:     "localhost",
		})
		if err != nil {
			log.Fatal(err)
		}

		book, err := dbConnection.GetBook(context.Background(), res.Book.Id)
		r.NoError(err)
		r.Equal(res.Book.Id, book.Id)
		r.Equal(res.Book.Author, book.Author)
		r.Equal(res.Book.Title, book.Title)
		r.Equal(res.Book.CreationTime, book.CreationTime)
	})

	t.Run("validation", func(t *testing.T) {
		// Verify that validation is performed by attempting to raise an
		// invalid argument via empty author.
		r := require.New(t)
		res, err := client.CreateBook(
			context.Background(),
			&books.CreateBookRequest{
				Book: &books.Book{
					Id:           ulid.Make().String(),
					Title:        ulid.Make().String(),
					Author:       "",
					CreationTime: timestamppb.New(time.Now()),
				},
				RequestId: ulid.Make().String(),
			},
		)
		r.Equal(codes.InvalidArgument, status.Code(err), "expected invalid argument")
		r.Zero(res)
	})
}
