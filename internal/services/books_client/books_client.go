package books_client

import (
	"log"

	books "github.com/celestebrant/library-of-books/books"
	"google.golang.org/grpc"
)

func MustNewClient() books.BooksClient {
	// Connect to server
	clientConnection, err := grpc.Dial("127.0.0.1:8089", grpc.WithInsecure())
	if err != nil {
		log.Panicf("failed to connect: %v", err)
	}
	defer clientConnection.Close()

	// Create client
	books_client := books.NewBooksClient(clientConnection)
	return books_client
}
