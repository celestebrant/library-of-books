# Docker-SQL demo

This project is for practicing the creation of a MySQL database in a docker container in Golang.

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