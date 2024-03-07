package booksservice

import (
	"fmt"
	"testing"
	"time"

	books "github.com/celestebrant/library-of-books/books"
	"github.com/oklog/ulid/v2"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestValidateAuthor(t *testing.T) {
	t.Parallel()

	t.Run("max length", func(t *testing.T) {
		r := require.New(t)
		author := `author with 255 characters: aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa`
		r.Len(author, 255)

		book := &books.Book{
			Author: author,
		}
		err := validateAuthor(book)
		r.NoError(err)
	})

	t.Run("max length + 1 returns error", func(t *testing.T) {
		r := require.New(t)
		author := `author with 256 characters: baaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa`
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
		title := `title with 255 characters: aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa`
		r.Len(title, 255)

		book := &books.Book{
			Title: title,
		}
		err := validateTitle(book)
		r.NoError(err)
	})

	t.Run("max length + 1 returns error", func(t *testing.T) {
		r := require.New(t)
		title := `title with 256 characters: baaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa`
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
		book := &books.Book{
			Id: ulid.Make().String(),
		}
		err := validateID(book)
		r.NoError(err)
	})

	t.Run("max length + 1 returns error", func(t *testing.T) {
		r := require.New(t)
		book := &books.Book{
			Id: fmt.Sprint(ulid.Make().String(), "a"),
		}
		err := validateID(book)
		var invalidIDError *InvalidIDError
		r.ErrorAs(err, &invalidIDError)
	})

	t.Run("empty returns error", func(t *testing.T) {
		r := require.New(t)
		book := &books.Book{
			Id: ``,
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
		book := &books.Book{
			Author:       `author1`,
			Title:        `title1`,
			Id:           ``,
			CreationTime: timestamppb.New(time.Now()),
		}

		err := Validate(book)
		var invalidIDError *InvalidIDError
		r.ErrorAs(err, &invalidIDError)
	})
}
