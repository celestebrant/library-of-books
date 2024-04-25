# Library Backend Application

## Introduction
This project is a backend application for a library management system. It currently allows users to  register and list books in the library database.

Planned features:
* ~~gRPC service for external transacting with the database~~ ✅
* ~~Ability to search for books by author, title, etc.~~ ✅
* Manage book borrowing and returns (checkout, return, overdue alerts, fines)
* User management (customer accounts, book bans)

## Getting started
This application requires a MySQL database and a gRPC server.

### Prerequisites
* Docker engine
* Go (with protocol compiler plugins, see [gRPC docs](https://grpc.io/docs/languages/go/quickstart/))
* Internet connection

### Database setup
1. Run `docker-compose up` to create and initialise the MySQL database using `docker-compose.yaml`.

### gRPC server setup
1. Generate gRPC code: `protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./books/books.proto`
2. Run the server: `go run ./cmd/server`

### Client setup
1. Start the server in a separate terminal.
2. Run the server: `go run ./cmd/client`

### Security note
The `docker-compose.yaml` file contains demo credentials for database access. This project is designed to be a demo. Do not use these credentials in production.

## Interacting with the database
This is not recommended for production use, but for demonstration purposes.
* Code example: Run `go run ./cmd/db_write_demo` to see a basic connection and write operation and fetch operation.
* Terminal access: `docker-compose exec database mysql -u user1 -p password1` (values for `-u` and `-p` are from `docker-compose.yaml`).

## Testing
The project uses unit and integration tests for code coverage. Run with `go clean -testcache; go test <path>`

### Unit tests
* `./internal/services/booksservice`
* `./storage`
* `./utils`

### Integration tests
* `go test ./tests`

### Planned test improvements
* gRPC calls and client setup
* Negative database write handling
* Addressing `TODO` comments in the codebase

# Supplementary information
## Database
The schema for table `books` is as follows, which resides in database `library`:
```
+---------------+--------------+------+-----+---------+-------+
| Field         | Type         | Null | Key | Default | Extra |
+---------------+--------------+------+-----+---------+-------+
| id            | varchar(30)  | NO   | PRI | NULL    |       |
| creation_time | timestamp    | YES  |     | NULL    |       |
| update_time   | timestamp    | YES  |     | NULL    |       |
| title         | varchar(255) | YES  |     | NULL    |       |
| author        | varchar(255) | YES  |     | NULL    |       |
+---------------+--------------+------+-----+---------+-------+
```
(Via queries: `USE library;` then `SHOW COLUMNS FROM books;`)

Handy commands:
* Show tables with `SHOW DATABASES;`
* Change database with `USE library;`
* Show tables with `SHOW TABLES;`. If you don't see the tables you expected, the program may be failing to find your volume.
