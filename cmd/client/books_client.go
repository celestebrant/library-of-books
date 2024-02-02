package main

import (
	"context"
	"log"
	"time"

	books "github.com/celestebrant/library-of-books/books"
	"github.com/celestebrant/library-of-books/internal/services/booksclient"
	"github.com/oklog/ulid/v2"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func main() {
	client, conn := booksclient.MustNewBooksClient("127.0.0.1:8089")
	defer conn.Close()

	res, err := client.CreateBook(
		context.Background(),
		&books.CreateBookRequest{
			Book: &books.Book{
				Id:           ulid.Make().String(),
				Title:        ulid.Make().String(),
				Author:       ulid.Make().String(),
				CreationTime: timestamppb.New(time.Now()),
			},
			RequestId: ulid.Make().String(),
		},
	)
	if err != nil {
		log.Fatalf("failed to call CreateBook: %v", err)
	}
	log.Printf("created book via CreateBook: %v", res)
}
