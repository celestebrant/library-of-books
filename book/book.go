package book

import "time"

type Book struct {
	ID int64
	Title string
	Author string
	CreationTime time.Time
}