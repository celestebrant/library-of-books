package main

import (
	"log"
	"net"

	books "github.com/celestebrant/library-of-books/books"
	"github.com/celestebrant/library-of-books/internal/services/books_service"
	"google.golang.org/grpc"
)

const address string = "127.0.0.1:8089"

func main() {
	// Create a network listener
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Create a new gRPC server registered with booksServer
	grpcServer := grpc.NewServer()
	books.RegisterBooksServer(grpcServer, &books_service.BooksServer{})

	// temporary address. Will replace with something dynamic through docker
	log.Printf("gRPC server listening on %s", address)

	// Connect the new server to the network listener
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve %v", err)
	}
}
