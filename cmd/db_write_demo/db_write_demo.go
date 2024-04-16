package main

import (
	"context"
	"fmt"
	"log"
	"time"

	storage "github.com/celestebrant/library-of-books/storage"
	"github.com/oklog/ulid/v2"
)

func main() {
	dbConnection, err := storage.NewMysqlStorage(storage.MysqlConfig{
		Username: "user1",
		Password: "password1",
		DBName:   "library",
		Port:     3306,
		Host:     "localhost", // this code will execute in machine, not container
	})
	if err != nil {
		log.Fatal(err)
	}

	author, title := "author1", "title1"
	book := &storage.Book{
		Id:           ulid.Make().String(),
		Title:        title,
		Author:       author,
		CreationTime: time.Now().UTC(),
	}
	if err = dbConnection.CreateBook(context.Background(), book); err != nil {
		log.Fatalf(`error encountered during CreateBook SQL operation: %v`, err)
	}

	log.Printf(`inserted record into "books" table: %v`, *book)

	books, err := dbConnection.ListBooks(context.Background(), author, title, 10, "")
	if err != nil {
		log.Fatalf(`error encountered during ListBooks SQL operation: %v`, err)
	}

	log.Printf(`fetched %d records via ListBooks from "books" table:`, len(books))
	if len(books) > 0 {
		for _, book := range books {
			fmt.Printf("- %v\n", book)
		}
	}
}
