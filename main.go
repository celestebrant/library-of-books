package main

import (
	"context"
	"log"
	"time"

	"github.com/celestebrant/library-of-books/book"
	storage "github.com/celestebrant/library-of-books/storage"
)

func main() {
	store, err := storage.NewMysqlStorage(storage.MysqlConfig{
		Username: "user1",
		Password: "password1",
		DBName:   "db",
		Port:     3306,
		Host:     "localhost", // this code will execute in machine, not container
	})

	if err != nil {
		log.Fatalf("failed to create mysql storage: %v", err)
	}

	log.Printf("mysql database store created: %v", store)

	b, err := store.CreateBook(context.Background(), &book.Book{
		CreationTime: time.Now(),
		Title:        "title1",
		Author:       "author1",
	})

	if err != nil {
		log.Fatalf("cannot insert book into table `books`: %v", err)
	}

	log.Printf("inserted book into table %v", b)
}
