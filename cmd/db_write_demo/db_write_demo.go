package main

import (
	"context"
	"log"
	"time"

	books "github.com/celestebrant/library-of-books/books"
	"google.golang.org/protobuf/types/known/timestamppb"

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

	b, err := dbConnection.CreateBook(context.Background(), &books.Book{
		CreationTime: timestamppb.New(time.Now()),
		Title:        "title1",
		Author:       "author1",
	})
	if err != nil {
		log.Fatalf("cannot insert book into table `books`: %v", err)
	}

	log.Printf("inserted book into table: %v", b)
}
