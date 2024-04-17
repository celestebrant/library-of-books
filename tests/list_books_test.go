package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/celestebrant/library-of-books/books"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/oklog/ulid/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestListBooks contains integration tests for the ListBooks service method and db.
func TestListBooks(t *testing.T) {
	// Prepare set up and tear down of server and client on different port.
	client, tearDown := setUpServerAndClient("127.0.0.1:8090")
	defer tearDown()

	t.Run("pagination", func(t *testing.T) {
		r := require.New(t)

		filter := ulid.Make().String()
		var pageSize int64 = 3

		// Make 5 books in total
		for i := 0; i < 5; i++ {
			id := fmt.Sprintf("%s_%d", filter, i)
			req := &books.CreateBookRequest{
				Book: &books.Book{
					Id:     id,
					Author: id,
					Title:  id,
				},
				RequestId: ulid.Make().String(),
			}
			_, err := client.CreateBook(context.Background(), req)
			r.NoError(err)
		}

		// First page should have 3 of 5 books
		res1, err := client.ListBooks(
			context.Background(),
			&books.ListBooksRequest{
				Author:   filter,
				Title:    filter,
				PageSize: pageSize,
			},
		)
		r.NoError(err)
		r.NotEmpty(res1)
		r.Len(res1.Books, 3)

		ignoreCreationTimeOpt := cmpopts.IgnoreFields(books.Book{}, "CreationTime")
		ignoreUnexportedOpt := cmpopts.IgnoreUnexported(books.Book{})

		for i, b := range res1.Books {
			id := fmt.Sprintf("%s_%d", filter, i)
			expected := &books.Book{
				Id:     id,
				Author: id,
				Title:  id,
			}
			if !cmp.Equal(expected, b, ignoreCreationTimeOpt, ignoreUnexportedOpt) {
				r.Fail(
					"books do not match",
					cmp.Diff(expected, b, ignoreCreationTimeOpt, ignoreUnexportedOpt),
				)
			}
		}
		r.NotEmpty(res1.NextPageToken)

		// Next page should have remaining 2 of 5 books
		res2, err := client.ListBooks(
			context.Background(),
			&books.ListBooksRequest{
				Author:        filter,
				Title:         filter,
				PageSize:      pageSize,
				NextPageToken: res1.NextPageToken,
			},
		)
		r.NoError(err)
		r.NotEmpty(res2)
		r.Len(res2.Books, 2)

		for i, b := range res2.Books {
			id := fmt.Sprintf("%s_%d", filter, i+3)
			expected := &books.Book{
				Id:     id,
				Author: id,
				Title:  id,
			}
			if !cmp.Equal(expected, b, ignoreCreationTimeOpt, ignoreUnexportedOpt) {
				r.Fail(
					"books do not match",
					cmp.Diff(expected, b, ignoreCreationTimeOpt, ignoreUnexportedOpt),
				)
			}
		}
		r.Empty(res2.NextPageToken)
	})

	t.Run("pagination empty next page if absolute total results is multiple of page size", func(t *testing.T) {
		// Create 2 books that satisfy the filters, and list with page size = 2
		r := require.New(t)

		filter := ulid.Make().String()
		var pageSize int64 = 2

		for i := 0; i < 2; i++ {
			id := fmt.Sprintf("%s_%d", filter, i)
			req := &books.CreateBookRequest{
				Book: &books.Book{
					Author: id,
					Title:  id,
				},
				RequestId: ulid.Make().String(),
			}

			_, err := client.CreateBook(context.Background(), req)
			r.NoError(err)
		}

		res1, err := client.ListBooks(
			context.Background(),
			&books.ListBooksRequest{
				Author:   filter,
				Title:    filter,
				PageSize: pageSize,
			},
		)
		r.NoError(err)
		r.NotEmpty(res1)
		r.Len(res1.Books, 2)
		r.NotEmpty(res1.NextPageToken)

		res2, err := client.ListBooks(
			context.Background(),
			&books.ListBooksRequest{
				Author:        filter,
				Title:         filter,
				PageSize:      pageSize,
				NextPageToken: res1.NextPageToken,
			},
		)
		r.NoError(err)
		r.NotEmpty(res2)
		r.Empty(res2.Books)
	})

	t.Run("no filters", func(t *testing.T) {
		// Create a book so that at least one exists, then perfom list call
		r := require.New(t)

		req := &books.CreateBookRequest{
			Book: &books.Book{
				Id:     ulid.Make().String(),
				Author: ulid.Make().String(),
				Title:  ulid.Make().String(),
			},
			RequestId: ulid.Make().String(),
		}
		_, err := client.CreateBook(context.Background(), req)
		r.NoError(err)

		res, err := client.ListBooks(
			context.Background(),
			&books.ListBooksRequest{
				PageSize: 5,
			},
		)
		r.NoError(err)
		r.NotEmpty(res)
		r.NotEmpty(res.Books)
	})

	t.Run("malformatted request returns error", func(t *testing.T) {
		r, a := require.New(t), assert.New(t)

		res, err := client.ListBooks(
			context.Background(),
			&books.ListBooksRequest{
				PageSize: 0,
			},
		)
		a.Error(err)
		r.Empty(res)
	})
}
