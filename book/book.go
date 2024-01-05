package book

import "time"

type Book struct {
	ID string
	Name string
	Author string
	CreationTime time.Time
}