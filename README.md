# Library of Books

## Introduction
This project is a representation of a back-end application for a library. Currently, it is possible to:
* Register a book to the library

(And more functionality will be added in the future.)

## Structure of the application
The library application comes in two main parts:
* A gRPC service, `Books`
* A MySQL database, `db`

### How it works
The `Books` service, defined in `./books/books.proto`, serves an interface upon which the client may affect records stored in the MySQL database table, `books`.

The schema for table `books` is as follows, which resides in database `db`:
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
(Via queries: `USE db;` then `SHOW COLUMNS FROM books;`)

## Initialising the database
It is possible to initialise the MySQL database in this project via a lightly manual process.

The relevant files are:
* `./docker-compose.yaml`, a configuration file used by Docker Compose to define and manage the containerised application.
* `./docker/dbinit/init.sql`, containing the schema.

Instructions:
1. Ensure you have the Docker engine installed and running. On Mac, open the Docker application.
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
1. Ensure the latest Go-generated proto files exist by running `make generate_grpc_code`. Not working? Try running the `protoc` command which you can find inside that file, and ensure your GOPATH is correct.
1. Run `go run ./server`
1. You are now ready to make calls.

## Using the deployed application
### Registering a book to the library
To interact with the gRPC interface, you need to set up a client. An example of this exists in `???`. Try it out with `???`.

This project is not yet complete and the following features shall be added in due course:
* ~~gRPC endpoints for external transacting with the database~~ ✅
* Ability for a customer to take out a book
* Ability for a customer to return a book
* Alert customers if the book due date has been reached
* Charge customers a fee for overdue books
* Ban customers from taking out books

### Interacting with the database directly
While this is not recommended for real-life use, it may be useful for exploratory or debugging purposes.

#### Write demo
A demo exists which creates a connection to the database and adds a book (writes a record to table `book`). You can run this with `go run ./db_write_demo`.

#### Connecting to the database via terminal
A demo username and password (`user1` and `password1`) is defined in `./docker-compose.yaml` intended for demo purposes only.

Run ```docker-compose exec database mysql -uuser1 -ppassword1``` to access the database via terminal.

Other handy database commands:
* Show tables with `SHOW DATABASES;`
* Change database with `USE <database name>;`, e.g. `USE db;`.
* Show tables with `SHOW TABLES;`. If you don't see the tables you expected, perhaps the program cannot find your volume.

## Run tests:
Tests exist for:
* Book validation, `go test ./validate_book`

Test cases to be added:
* gRPC calls (and client set up)
* Negative db write handling