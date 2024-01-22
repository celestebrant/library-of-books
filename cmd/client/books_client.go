package main

import (
	"context"
	"log"
	"time"

	books "github.com/celestebrant/library-of-books/books"
	"github.com/oklog/ulid/v2"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func main() {
	// Connect to server
	clientConnection, err := grpc.Dial("127.0.0.1:8089", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer clientConnection.Close()

	// Create client
	books_client := books.NewBooksClient(clientConnection)

	// Call gRPC method
	res, err := books_client.CreateBook(
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
	log.Printf("Called CreateBook: %v", res)
}
