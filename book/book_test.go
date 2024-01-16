package book

import (
	"fmt"
	"testing"
	"time"

	"github.com/oklog/ulid/v2"
	"github.com/stretchr/testify/require"
)

func TestValidateAuthor(t *testing.T) {
	t.Parallel()

	t.Run("max length", func(t *testing.T) {
		r := require.New(t)
		author := `author with 255 characters: aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa`
		r.Len(author, 255)

		book := Book{
			Author: author,
		}
		err := book.validateAuthor()
		r.NoError(err)
	})

	t.Run("max length + 1 returns error", func(t *testing.T) {
		r := require.New(t)
		author := `author with 256 characters: baaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa`
		r.Len(author, 256)

		book := Book{
			Author: author,
		}
		err := book.validateAuthor()
		var invalidAuthorError *InvalidAuthorError
		r.ErrorAs(err, &invalidAuthorError)
	})

	t.Run("empty returns error", func(t *testing.T) {
		r := require.New(t)
		book := Book{
			Author: ``,
		}
		err := book.validateAuthor()
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

		book := Book{
			Title: title,
		}
		err := book.validateTitle()
		r.NoError(err)
	})

	t.Run("max length + 1 returns error", func(t *testing.T) {
		r := require.New(t)
		title := `title with 256 characters: baaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa`
		r.Len(title, 256)

		book := Book{
			Title: title,
		}
		err := book.validateTitle()
		var invalidTitleError *InvalidTitleError
		r.ErrorAs(err, &invalidTitleError)
	})

	t.Run("empty returns error", func(t *testing.T) {
		r := require.New(t)
		book := Book{
			Title: ``,
		}
		err := book.validateTitle()
		var invalidTitleError *InvalidTitleError
		r.ErrorAs(err, &invalidTitleError)
	})
}

func TestValidateID(t *testing.T) {
	t.Run("max length", func(t *testing.T) {
		r := require.New(t)
		book := Book{
			Id: ulid.Make().String(),
		}
		err := book.validateID()
		r.NoError(err)
	})

	t.Run("max length + 1 returns error", func(t *testing.T) {
		r := require.New(t)
		book := Book{
			Id: fmt.Sprint(ulid.Make().String(), "a"),
		}
		err := book.validateID()
		var invalidIDError *InvalidIDError
		r.ErrorAs(err, &invalidIDError)
	})

	t.Run("empty returns error", func(t *testing.T) {
		r := require.New(t)
		book := Book{
			Id: ``,
		}
		err := book.validateID()
		var invalidIDError *InvalidIDError
		r.ErrorAs(err, &invalidIDError)
	})
}

func TestValidateCreationTime(t *testing.T) {
	t.Run("now", func(t *testing.T) {
		r := require.New(t)
		book := Book{
			CreationTime: time.Now(),
		}
		err := book.validateCreationTime()
		r.NoError(err)
	})

	t.Run("zero value returns error", func(t *testing.T) {
		r := require.New(t)
		book := Book{
			CreationTime: time.Time{},
		}

		err := book.validateCreationTime()
		var invalidCreationTimeError *InvalidCreationTimeError
		r.ErrorAs(err, &invalidCreationTimeError)
	})
}

func TestValidate(t *testing.T) {
	t.Run("positive returns no error", func(t *testing.T) {
		r := require.New(t)
		book := Book{
			Author:       `author1`,
			Title:        `title1`,
			Id:           ulid.Make().String(),
			CreationTime: time.Now(),
		}
		err := book.Validate()
		r.NoError(err)
	})

	t.Run("validates author", func(t *testing.T) {
		r := require.New(t)
		book := Book{
			Author:       ``,
			Title:        `title1`,
			Id:           ulid.Make().String(),
			CreationTime: time.Now(),
		}

		err := book.Validate()
		var invalidAuthorError *InvalidAuthorError
		r.ErrorAs(err, &invalidAuthorError)
	})

	t.Run("validates title", func(t *testing.T) {
		r := require.New(t)
		book := Book{
			Author:       `author1`,
			Title:        ``,
			Id:           ulid.Make().String(),
			CreationTime: time.Now(),
		}

		err := book.Validate()
		var invalidTitleError *InvalidTitleError
		r.ErrorAs(err, &invalidTitleError)
	})

	t.Run("validates ID", func(t *testing.T) {
		r := require.New(t)
		book := Book{
			Author:       `author1`,
			Title:        `title1`,
			Id:           ``,
			CreationTime: time.Now(),
		}

		err := book.Validate()
		var invalidIDError *InvalidIDError
		r.ErrorAs(err, &invalidIDError)
	})

	t.Run("validates creation time", func(t *testing.T) {
		r := require.New(t)
		book := Book{
			Author:       `author1`,
			Title:        `title1`,
			Id:           ulid.Make().String(),
			CreationTime: time.Time{},
		}

		err := book.Validate()
		var invalidCreationTimeError *InvalidCreationTimeError
		r.ErrorAs(err, &invalidCreationTimeError)
	})
}
