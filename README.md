# Library of Books

## Introduction
This project is a representation of a back-end application for a library. Currently, it is possible to:
* Register a book to the library

This project is not yet complete and the following features shall be added in due course:
* ~~gRPC service for external transacting with the database~~ ✅
* Ability to search for books by author, title, etc.
* Ability for a customer to take out a book
* Ability for a customer to return a book
* Alert customers if the book due date has been reached
* Charge customers a fee for overdue books
* Ban customers from taking out books

Known issues and limitations:
* `make generate_grpc_code` fails. Current workaround is to run the `protoc` command directly into terminal.
* All packages in this project are publically visible.

## Structure of the application
The library application comes in two main parts, in order of required deployment for use:
1. A MySQL database, `library`
1. A gRPC service, `Books`

After the application is running fully, you may create a client which can interact with the `books` service. See section "Setting up a client" for more on this.

### How it works
The `Books` service, defined in `./books/books.proto` allows the client to perform actions such as creating books. It serves an interface upon which a client may affect records stored in the `books` table of the `library` database.

The schema for table `books` is as follows, which resides in database `library`:
```
+---------------+--------------+------+-----+---------+-------+
| Field         | Type         | Null | Key | Default | Extra |
+---------------+--------------+------+-----+---------+-------+
| id            | varchar(26)  | NO   | PRI | NULL    |       |
| creation_time | timestamp    | YES  |     | NULL    |       |
| update_time   | timestamp    | YES  |     | NULL    |       |
| title         | varchar(255) | YES  |     | NULL    |       |
| author        | varchar(255) | YES  |     | NULL    |       |
+---------------+--------------+------+-----+---------+-------+
```
(Via queries: `USE library;` then `SHOW COLUMNS FROM books;`)

## Initialising the database
It is possible to initialise the MySQL database in this project via `docker-compose`.

The relevant files are:
* `./docker-compose.yaml`, a configuration file used by Docker Compose to define and manage the containerised application.
* `./docker/dbinit/init.sql`, containing the schema.

Instructions:
1. Ensure you have the Docker engine installed and running. On Mac, open the Docker application. Also ensure you are connected to the internet.
1. In a terminal window, run `docker-compose up`. This reads `docker-compose.yaml` and creates a container ready for the MySQL database.
1. If this is left running, then the database is ready.

If you make any changes to the database structure and you don't care about the records held, then you will need to reinitialise the database. This can be done by:

1. `CTRL + C` the window which has `docker-compose up` running
1. Run `docker-compose down`
1. Implement your changes
1. Run `docker-compose up --force-recreate`

Alternatively, close the Containers, Images and Volumes via the Docker engine directly. (This works as a suitable last resort.)

## Setting up the gRPC server
It is currently only possible to set up the `Books` gRPC server locally.

Instructions:
1. Ensure the latest Go-generated proto files exist by running `make generate_grpc_code`. You will need Go plugins for the protocol compiler. See the [gRPC docs](https://grpc.io/docs/languages/go/quickstart/) on how set this up.
1. Run `go run ./cmd/server`.
1. You are now ready to make calls. Your terminal should look like it is hanging.

## Setting up a client
To interact with the gRPC server, you need to set up a client. An example client exists which you can try running with `go run ./cmd/client`, once a server is already running (in a separate terminal window).

## Interacting with the database directly
While this is not recommended for real-life use, it may be useful for exploratory or debugging purposes.

### Querying via code
A demo exists which creates a connection to the database and adds a book (writes a record to table `book`). You can run this with `go run ./cmd/db_write_demo`.

### Querying via terminal
A demo username and password (`user1` and `password1`) is defined in `./docker-compose.yaml` intended for demo purposes only.

Run ```docker-compose exec database mysql -uuser1 -ppassword1``` to access the database via terminal.

Other handy database commands:
* Show tables with `SHOW DATABASES;`
* Change database with `USE <database name>;`, e.g. `USE library;`.
* Show tables with `SHOW TABLES;`. If you don't see the tables you expected, perhaps the program cannot find your volume.

## Tests
Packages `testing` (internal) and `stretchr` (third party) are used in the tests.
All test files look like `*_test.go`. Unit tests can be found within packages containing logic. Higher level tests can be found in `./tests/`.

Coverage exists for:
* Unit tests for book validation, `go test ./internal/services/booksservice`
* Integration tests, `go test ./tests`

Test cases to be added:
* gRPC calls (and client set up)
* Negative db write handling
* `TODO` comments in the codebase
