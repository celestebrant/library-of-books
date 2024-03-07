package main

import (
	"context"
	"log"

	storage "github.com/celestebrant/library-of-books/storage"
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

	book := &storage.Book{
		Title:  "title1",
		Author: "author1",
	}
	if err = dbConnection.CreateBook(context.Background(), book); err != nil {
		log.Fatalf("cannot insert into table `books`: %v", err)
	}

	log.Printf("inserted book into table: %v", book)

	// TODO: GetBook
}
