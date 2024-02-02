package booksclient

import (
	"log"

	books "github.com/celestebrant/library-of-books/books"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// MustNewBooksClient creates and returns a new books client and its client connection,
// or panics if an error is encountered. It is recommended that any client connection is
// subsequently closed with a deferred *grpc.ClientConn.Close().
func MustNewBooksClient(address string) (books.BooksClient, *grpc.ClientConn) {
	// Connect to server
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Panicf("failed to create books gRPC client: %v", err)
	}

	// Create client
	client := books.NewBooksClient(conn)
	return client, conn
}
