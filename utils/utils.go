package utils

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"strconv"
)

// Offset returns the offset for a given page token.
func Offset(pageToken string) (int64, error) {
	if pageToken == "" {
		return 0, nil
	}

	pageTokenBytes, err := base64.StdEncoding.DecodeString(pageToken)
	if err != nil {
		return 0, fmt.Errorf(`failed to decode page token "%s": %w`, pageToken, err)
	}

	offset, err := strconv.ParseInt(string(pageTokenBytes), 10, 64)
	if err != nil {
		return 0, fmt.Errorf(`failed to parse decoded page token into int64: %w`, err)
	}

	return offset, nil
}

// NextPageToken returns the next page token for a given previous page token by summing its offset with page size.
func NextPageToken(prevPageToken string, pageSize int64) (string, error) {
	prevOffset, err := Offset(prevPageToken)
	if err != nil {
		return "", fmt.Errorf(`failed to find offset of previous page token: %w`, err)
	}

	return pageToken(prevOffset + pageSize), nil
}

// pageToken returns the page token for a given offset.
func pageToken(offset int64) string {
	if offset == 0 {
		return ""
	}
	return base64.StdEncoding.EncodeToString([]byte(strconv.FormatInt(offset, 10)))
}

// StringWithLength produces a string with the length specified, like "aaaaa".
func StringWithLength(length int) string {
	return string(bytes.Repeat([]byte{byte('a')}, length))
}
