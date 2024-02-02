package booksservice

import (
	"context"
	"log"
	"net"
	"sync"

	books "github.com/celestebrant/library-of-books/books"
	"github.com/celestebrant/library-of-books/storage"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// MustNewBooksServer creates a new BooksServer in a goroutine, and panics if setup fails.
// Returns the books gRPC server and its network listener, and subsequence closures
// should be deferred with *grpc.Server.Stop() and net.Listener.Close().
func MustNewBooksServer(address string) (*grpc.Server, net.Listener, *sync.WaitGroup) {
	var wg sync.WaitGroup

	// Create a network listener
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Panicf("failed to listen: %v", err)
	}

	// Create a new database connection that the books server can use to write to the db
	dbConn, err := storage.NewMysqlStorage(storage.MysqlConfig{
		Username: "user1",
		Password: "password1",
		DBName:   "library",
		Port:     3306,
		Host:     "localhost", // this code will execute in machine, not container
	})
	if err != nil {
		log.Panic(err)
	}

	// Create a new gRPC server registered with booksServer
	grpcServer := grpc.NewServer()
	books.RegisterBooksServer(grpcServer, &BooksServer{
		MysqlStorage: &dbConn,
	})
	log.Printf("gRPC server books listening on %s", address)

	// Connect the new server to the network listener in a goroutine
	wg.Add(1)
	go func() {
		defer wg.Done()
		err = grpcServer.Serve(lis)
		if err != nil {
			log.Panicf("failed to serve: %v", err)
		}
	}()

	return grpcServer, lis, &wg
}

// StopBooksServer stops server, closes lis and waits until wg is zero.
func StopBooksServer(server *grpc.Server, lis net.Listener, wg *sync.WaitGroup) {
	server.Stop()
	lis.Close()
	wg.Wait()
}

// BooksServer represents the books service and implements storage.MysqlStorage
// to enable database connections.
type BooksServer struct {
	books.UnimplementedBooksServer
	*storage.MysqlStorage
}

// CreateBook semantically validates the Book contained in req and creates a record of book
// in the database. Returns the created book in CreateBookResponse.
func (s *BooksServer) CreateBook(
	ctx context.Context, req *books.CreateBookRequest,
) (*books.CreateBookResponse, error) {
	err := Validate(req.Book)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	// TODO: decide if returned Book should be used in res. If no, add tests to compare res.Book with db-returned book
	_, err = s.MysqlStorage.CreateBook(ctx, req.Book)

	res := &books.CreateBookResponse{Book: req.Book}
	return res, nil
}