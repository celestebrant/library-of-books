package booksservice

import (
	"testing"

	"github.com/celestebrant/library-of-books/books"
	"github.com/celestebrant/library-of-books/utils"
	"github.com/stretchr/testify/require"
)

// newValidCreateBookRequest returns a new valid CreateBookRequest where fields have
// their maximum accepted length.
func newValidCreateBookRequest() *books.CreateBookRequest {
	return &books.CreateBookRequest{
		RequestId: utils.StringWithLength(requestIDMaxLength),
		Book: &books.Book{
			Id:     utils.StringWithLength(idMaxLength),
			Author: utils.StringWithLength(authorMaxLength),
			Title:  utils.StringWithLength(titleMaxLength),
		},
	}
}

func TestValidateCreateBookRequest(t *testing.T) {
	t.Parallel()

	t.Run("valid", func(t *testing.T) {
		r := require.New(t)
		req := newValidCreateBookRequest()
		err := ValidateCreateBookRequest(req)
		r.NoError(err)
	})

	t.Run("omitted request ID returns error", func(t *testing.T) {
		r := require.New(t)

		req := newValidCreateBookRequest()
		req.RequestId = ""

		err := ValidateCreateBookRequest(req)
		expectedErr := ValidationError{
			"request_id",
			"must not be empty",
		}
		r.EqualError(err, expectedErr.Error())
	})

	t.Run("request ID max length + 1 returns error", func(t *testing.T) {
		r := require.New(t)

		req := newValidCreateBookRequest()
		req.RequestId = utils.StringWithLength(31)

		err := ValidateCreateBookRequest(req)
		expectedErr := ValidationError{
			"request_id",
			"must not exceed 30 characters",
		}
		r.EqualError(err, expectedErr.Error())
	})

	t.Run("omitted author returns error", func(t *testing.T) {
		r := require.New(t)

		req := newValidCreateBookRequest()
		req.Book.Author = ""

		err := ValidateCreateBookRequest(req)
		expectedErr := ValidationError{
			"author",
			"must not be empty",
		}
		r.EqualError(err, expectedErr.Error())
	})

	t.Run("author max length + 1 returns error", func(t *testing.T) {
		r := require.New(t)

		req := newValidCreateBookRequest()
		req.Book.Author = utils.StringWithLength(256)

		err := ValidateCreateBookRequest(req)
		expectedErr := ValidationError{
			"author",
			"must not exceed 255 characters",
		}
		r.EqualError(err, expectedErr.Error())
	})

	t.Run("omitted title returns error", func(t *testing.T) {
		r := require.New(t)

		req := newValidCreateBookRequest()
		req.Book.Title = ""

		err := ValidateCreateBookRequest(req)
		expectedErr := ValidationError{
			"title",
			"must not be empty",
		}
		r.EqualError(err, expectedErr.Error())
	})

	t.Run("title max length + 1 returns error", func(t *testing.T) {
		r := require.New(t)

		req := newValidCreateBookRequest()
		req.Book.Title = utils.StringWithLength(256)

		err := ValidateCreateBookRequest(req)
		expectedErr := ValidationError{
			"title",
			"must not exceed 255 characters",
		}
		r.EqualError(err, expectedErr.Error())
	})

	t.Run("omitted id is accepted", func(t *testing.T) {
		r := require.New(t)

		req := newValidCreateBookRequest()
		req.Book.Id = ""

		err := ValidateCreateBookRequest(req)
		r.NoError(err)
	})

	t.Run("id max length + 1 returns error", func(t *testing.T) {
		r := require.New(t)

		req := newValidCreateBookRequest()
		req.Book.Id = utils.StringWithLength(idMaxLength + 1)

		err := ValidateCreateBookRequest(req)
		expectedErr := ValidationError{
			"id",
			"must not exceed 30 characters",
		}
		r.EqualError(err, expectedErr.Error())
	})
}

func TestValidateListBooksRequest(t *testing.T) {
	t.Parallel()

	t.Run("valid lower boundary", func(t *testing.T) {
		r := require.New(t)

		req := &books.ListBooksRequest{
			PageSize: 1,
		}
		err := ValidateListBooksRequest(req)
		r.NoError(err)
	})

	t.Run("valid upper boundary", func(t *testing.T) {
		r := require.New(t)

		req := &books.ListBooksRequest{
			PageSize: 50,
		}
		err := ValidateListBooksRequest(req)
		r.NoError(err)
	})

	t.Run("invalid lower boundary", func(t *testing.T) {
		r := require.New(t)

		req := &books.ListBooksRequest{
			PageSize: 0,
		}
		err := ValidateListBooksRequest(req)
		expectedErr := ValidationError{
			"page_size",
			"must be greater than zero and not exceed 50",
		}
		r.EqualError(err, expectedErr.Error())
	})

	t.Run("invalid lower boundary", func(t *testing.T) {
		r := require.New(t)

		req := &books.ListBooksRequest{
			PageSize: 51,
		}
		err := ValidateListBooksRequest(req)
		expectedErr := ValidationError{
			"page_size",
			"must be greater than zero and not exceed 50",
		}
		r.EqualError(err, expectedErr.Error())
	})
}
