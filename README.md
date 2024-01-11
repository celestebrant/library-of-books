# Library of Books

This project a representation of a back-end application for a library, for demo purposes. It contains an initialisation of a MySQL database into a docker container.

Currently, it is possible to:
* Register a book to the library

This project is not yet complete and the following features shall be added in due course:
* gRPC endpoints for external transacting with the database
* Ability for a customer to take out a book
* Ability for a customer to return a book
* Alert customers if the book due date has been reached
* Charge customers a fee for overdue books
* Ban customers from taking out books

## Set up instructions
1. Ensure the docker engine is running on your machine. In Mac, open the Docker application.
1. In a terminal window, run `docker-compose up`. This reads `docker-compose.yaml` and creates a container ready for the MySQL database.
1. You are now ready.

## Usage instructions
You can make calls to the database. A demo exists in `main.go` which creates a connection to the database and adds a book. You can run this with `go run main.go`.

## Database handling
If at any point you need to amend the database:
1. In the terminal window that is running docker-compose, hit `CTRL + C`.
1. Run `docker-compose down`.
1. Make amendments to code.
1. Run `docker-compose up --force-recreate`
1. Then execute your code.

### Connecting to the database via terminal
```docker-compose exec database mysql -uuser1 -ppassword1```
This uses a demo username of `user1` and password of `password1`. You should now be able to query the database directly.

Other handy database commands
* Show tables with `SHOW DATABASES;`
* Change database with `USE <database name>;`, e.g. `USE db;`.
* Show tables with `SHOW TABLES;`. If you don't see the tables you expected, perhaps the program cannot find your volume.

Run tests:
* `go test ./book`