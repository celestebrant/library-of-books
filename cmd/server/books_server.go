package main

import (
	"log"
	"net"

	books "github.com/celestebrant/library-of-books/books"
	"github.com/celestebrant/library-of-books/internal/services/booksservice"
	"github.com/celestebrant/library-of-books/storage"
	"google.golang.org/grpc"
)

const address string = "127.0.0.1:8089"

func main() {
	// Create a network listener
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Create a new database connection that the books server can use to write to the db
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

	// Create a new gRPC server registered with booksServer
	grpcServer := grpc.NewServer()
	books.RegisterBooksServer(grpcServer, &booksservice.BooksServer{
		MysqlStorage: &dbConnection,
	})
	log.Printf("gRPC server listening on %s", address)

	// Connect the new server to the network listener
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve %v", err)
	}
}
