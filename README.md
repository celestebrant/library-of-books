# Library of Books

This project is a back-end application for a library. It contains an initialisation of a MySQL database into a docker container.

It is possible to:
* Register a book to the library

This project is not yet complete and the following features shall be added in due course:
* Ability for a customer to take out a book
* Ability for a customer to return a book
* Alert customers if the book due date has been reached
* Charge customers a fee for overdue books
* Ban customers from taking out books

Instructions
1. Ensure the docker engine is running on your machine. In Mac, open the Docker application.
1. In a terminal window, run `docker-compose up`. This reads `docker-compose.yaml` and creates a container ready for the MySQL database.
1. Execute your code by entering the following in another termincal window: `go run main.go`.

If at any point you need to amend the code and try running again:
1. In the terminal window that is running docker-compose, hit `CTRL + C`.
1. Run `docker-compose down`.
1. Make amendments to code.
1. Run `docker-compose up --force-recreate`
1. Then execute your code.

To connect to the database:
1. In a terminal window, run `docker-compose exec database mysql -uuser1 -ppassword1`. You should now be able to query the database.

Other database commands
* Show tables with `SHOW DATABASES;`
* Change database with `USE <database name>;`, e.g. `USE db;`.
* Show tables with `SHOW TABLES;`. If you don't see the tables you expected, perhaps the program cannot find your volume.

Run tests:
* `go test ./book`