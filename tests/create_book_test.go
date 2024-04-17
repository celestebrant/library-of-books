package tests

import (
	"context"
	"database/sql"
	"log"
	"sync"
	"testing"
	"time"

	"github.com/celestebrant/library-of-books/books"
	"github.com/celestebrant/library-of-books/internal/services/booksclient"
	"github.com/celestebrant/library-of-books/internal/services/booksservice"
	"github.com/celestebrant/library-of-books/storage"
	"github.com/oklog/ulid/v2"
	"github.com/stretchr/testify/assert"
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

// TestCreateBook contains integration tests for the CreateBook endpoint and db.
func TestCreateBook(t *testing.T) {
	// Prepare set up and tear down of server and client on different port.
	client, tearDown := setUpServerAndClient("127.0.0.1:8090")
	defer tearDown()

	t.Run("request with mandatory request fields only writes to db", func(t *testing.T) {
		r, a := require.New(t), assert.New(t)
		testStartTime := time.Now()
		req := &books.CreateBookRequest{
			Book: &books.Book{
				Title:  ulid.Make().String(),
				Author: ulid.Make().String(),
			},
			RequestId: ulid.Make().String(),
		}
		res, err := client.CreateBook(context.Background(), req)
		r.NoError(err)
		r.NotEmpty(res)

		// Validate RPC response against request
		a.Equal(req.Book.Title, res.Book.Title)
		a.Equal(req.Book.Author, res.Book.Author)
		a.Truef(
			res.Book.CreationTime.AsTime().After(testStartTime),
			"expected response book to have creation time %v after test start time %v",
			res.Book.CreationTime.AsTime(),
			testStartTime,
		)
		a.NotEmpty(res.Book.Id)

		// Validate db record against request
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

		a.NotEmpty(book.Id)
		a.Equal(req.Book.Author, book.Author)
		a.Equal(req.Book.Title, book.Title)
		a.True(book.CreationTime.After(testStartTime))
	})

	t.Run("request with all fields populated writes to db", func(t *testing.T) {
		r, a := require.New(t), assert.New(t)

		now := time.Now()
		req := &books.CreateBookRequest{
			Book: &books.Book{
				Id:           ulid.Make().String(),
				Title:        ulid.Make().String(),
				Author:       ulid.Make().String(),
				CreationTime: timestamppb.New(now),
			},
			RequestId: ulid.Make().String(),
		}
		res, err := client.CreateBook(context.Background(), req)
		r.NoError(err)
		r.NotEmpty(res)

		// Validate RPC response against request
		a.Equal(req.Book.Id, res.Book.Id)
		a.Equal(req.Book.Title, res.Book.Title)
		a.Equal(req.Book.Author, res.Book.Author)
		a.Equal(req.Book.CreationTime.AsTime(), res.Book.CreationTime.AsTime())

		// Validate db record against request
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

		r.Equal(req.Book.Id, book.Id)
		r.Equal(req.Book.Author, book.Author)
		r.Equal(req.Book.Title, book.Title)
		r.Equal(req.Book.CreationTime.AsTime(), book.CreationTime)
	})

	t.Run("validation error does not write to db", func(t *testing.T) {
		// Verify that validation is performed by attempting to raise an
		// invalid argument via empty author.
		r, a := require.New(t), assert.New(t)
		req := &books.CreateBookRequest{
			Book: &books.Book{
				Id:           ulid.Make().String(),
				Title:        ulid.Make().String(),
				Author:       "",
				CreationTime: timestamppb.New(time.Now()),
			},
			RequestId: ulid.Make().String(),
		}
		res, err := client.CreateBook(context.Background(), req)
		a.Equal(codes.InvalidArgument, status.Code(err), "expected invalid argument")
		r.Zero(res)

		// Verify no db record exists
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

		book, err := dbConnection.GetBook(context.Background(), req.Book.Id)
		a.ErrorIs(err, sql.ErrNoRows, "expected sql.ErrNoRows type error if no record found")
		a.Empty(book)
	})
}
