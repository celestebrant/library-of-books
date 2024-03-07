package storage

import (
	"testing"
	"time"

	"github.com/celestebrant/library-of-books/books"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestNewBook(t *testing.T) {
	t.Parallel()

	t.Run("all fields populated", func(t *testing.T) {
		r := require.New(t)
		creationTime := time.Now().UTC()
		req := &books.CreateBookRequest{
			Book: &books.Book{
				Id:           "hello",
				Title:        "Some title",
				Author:       "Some author",
				CreationTime: timestamppb.New(creationTime),
			},
		}
		book := NewBook(req)
		expected := Book{
			Id:           "hello",
			Title:        "Some title",
			Author:       "Some author",
			CreationTime: creationTime,
		}
		r.Equal(expected, *book)
	})

	t.Run("populates empty Id and CreationTime", func(t *testing.T) {
		a := assert.New(t)
		r := require.New(t)
		testStartTime := time.Now()
		req := &books.CreateBookRequest{
			Book: &books.Book{},
		}
		r.Equal(
			time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
			req.Book.CreationTime.AsTime(),
			"unexpected default value of CreateBookRequest.Book.CreationTimestamp",
		)
		book := NewBook(req)
		a.NotEmpty(book.Id, "expected Id to be populated")
		a.True(
			book.CreationTime.UTC().After(testStartTime.UTC()),
			"expected book.CreationTime %v to be after earlier 'now' time %v",
			book.CreationTime,
			testStartTime.UTC(),
		)
	})
}
