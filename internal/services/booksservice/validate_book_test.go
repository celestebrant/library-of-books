package booksservice

import (
	"testing"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/celestebrant/library-of-books/books"
	"github.com/celestebrant/library-of-books/utils"
	"github.com/oklog/ulid/v2"
	"github.com/stretchr/testify/require"
)

func TestValidateAuthor(t *testing.T) {
	t.Parallel()

	t.Run("max length", func(t *testing.T) {
		r := require.New(t)
		author := utils.StringWithLength(authorMaxLength)
		r.Len(author, 255)

		book := &books.Book{
			Author: author,
		}
		err := validateAuthor(book)
		r.NoError(err)
	})

	t.Run("max length + 1 returns error", func(t *testing.T) {
		r := require.New(t)
		author := utils.StringWithLength(authorMaxLength + 1)
		r.Len(author, 256)

		book := &books.Book{
			Author: author,
		}
		err := validateAuthor(book)
		var invalidAuthorError *InvalidAuthorError
		r.ErrorAs(err, &invalidAuthorError)
	})

	t.Run("empty returns error", func(t *testing.T) {
		r := require.New(t)
		book := &books.Book{
			Author: ``,
		}
		err := validateAuthor(book)
		var invalidAuthorError *InvalidAuthorError
		r.ErrorAs(err, &invalidAuthorError)
	})
}

func TestValidateTitle(t *testing.T) {
	t.Parallel()

	t.Run("max length", func(t *testing.T) {
		r := require.New(t)
		title := utils.StringWithLength(titleMaxLength)
		r.Len(title, 255)

		book := &books.Book{
			Title: title,
		}
		err := validateTitle(book)
		r.NoError(err)
	})

	t.Run("max length + 1 returns error", func(t *testing.T) {
		r := require.New(t)
		title := utils.StringWithLength(titleMaxLength + 1)
		r.Len(title, 256)

		book := &books.Book{
			Title: title,
		}
		err := validateTitle(book)
		var invalidTitleError *InvalidTitleError
		r.ErrorAs(err, &invalidTitleError)
	})

	t.Run("empty returns error", func(t *testing.T) {
		r := require.New(t)
		book := &books.Book{
			Title: ``,
		}
		err := validateTitle(book)
		var invalidTitleError *InvalidTitleError
		r.ErrorAs(err, &invalidTitleError)
	})
}

func TestValidateID(t *testing.T) {
	t.Run("max length", func(t *testing.T) {
		r := require.New(t)
		maxLenStr := utils.StringWithLength(idMaxLength)
		book := &books.Book{
			Id: maxLenStr,
		}
		err := validateID(book)
		r.NoError(err)
	})

	t.Run("max length + 1 returns error", func(t *testing.T) {
		r := require.New(t)
		exceedMaxLen := utils.StringWithLength(idMaxLength + 1)
		book := &books.Book{
			Id: exceedMaxLen,
		}
		err := validateID(book)
		var invalidIDError *InvalidIDError
		r.ErrorAs(err, &invalidIDError)
	})
}

func TestValidate(t *testing.T) {
	t.Run("positive returns no error", func(t *testing.T) {
		r := require.New(t)
		book := &books.Book{
			Author:       `author1`,
			Title:        `title1`,
			Id:           ulid.Make().String(),
			CreationTime: timestamppb.New(time.Now()),
		}
		err := Validate(book)
		r.NoError(err)
	})

	t.Run("validates author", func(t *testing.T) {
		r := require.New(t)
		book := &books.Book{
			Author:       ``,
			Title:        `title1`,
			Id:           ulid.Make().String(),
			CreationTime: timestamppb.New(time.Now()),
		}

		err := Validate(book)
		var invalidAuthorError *InvalidAuthorError
		r.ErrorAs(err, &invalidAuthorError)
	})

	t.Run("validates title", func(t *testing.T) {
		r := require.New(t)
		book := &books.Book{
			Author:       `author1`,
			Title:        ``,
			Id:           ulid.Make().String(),
			CreationTime: timestamppb.New(time.Now()),
		}

		err := Validate(book)
		var invalidTitleError *InvalidTitleError
		r.ErrorAs(err, &invalidTitleError)
	})

	t.Run("validates ID", func(t *testing.T) {
		r := require.New(t)
		exceedMaxLen := utils.StringWithLength(idMaxLength + 1)
		book := &books.Book{
			Author:       `author1`,
			Title:        `title1`,
			Id:           exceedMaxLen,
			CreationTime: timestamppb.New(time.Now()),
		}
		err := Validate(book)
		var invalidIDError *InvalidIDError
		r.ErrorAs(err, &invalidIDError)
	})
}

func TestValidateListBooksRequest(t *testing.T) {
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
		var invalidPageSizeError *InvalidPageSizeError
		r.ErrorAs(err, &invalidPageSizeError)
	})

	t.Run("invalid lower boundary", func(t *testing.T) {
		r := require.New(t)

		req := &books.ListBooksRequest{
			PageSize: 51,
		}
		err := ValidateListBooksRequest(req)
		var invalidPageSizeError *InvalidPageSizeError
		r.ErrorAs(err, &invalidPageSizeError)
	})
}
