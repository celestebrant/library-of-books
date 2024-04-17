package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestZeroOffsetPageToken(t *testing.T) {
	r := require.New(t)

	offset, err := Offset("")
	r.NoError(err)
	r.Equal(int64(0), offset)

	pageToken := pageToken(offset)
	r.Equal("", pageToken)
}

func TestTenOffsetPageToken(t *testing.T) {
	r := require.New(t)

	pageToken := pageToken(10)
	r.Equal("MTA=", pageToken)

	offset, err := Offset(pageToken)
	r.NoError(err)
	r.Equal(int64(10), offset)
}

func TestNextPageToken(t *testing.T) {
	r := require.New(t)

	pageSize := 10
	firstPageToken := ""
	secondPageToken, err := NextPageToken(firstPageToken, int64(pageSize))
	r.NoError(err)
	r.Equal("MTA=", secondPageToken)
}

func TestStringWithLength(t *testing.T) {
	r := require.New(t)
	str := StringWithLength(10)
	r.Len(str, 10)
}
