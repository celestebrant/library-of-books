package storage

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" // blank import runs init
)

type MysqlStorage struct {
	db *sql.DB
}

type MysqlConfig struct {
	Username string
	Password string
	DBName   string
	Port     uint
	Host     string
}

// NewMysqlStorage opens a new connection to the DB via the mysql driver.
func NewMysqlStorage(conf MysqlConfig) (MysqlStorage, error) {
	// username:password@protocol(address)/dbname?param=value
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", conf.Username, conf.Password, conf.Host, conf.Port, conf.DBName)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return MysqlStorage{}, fmt.Errorf("cannot validate open MySQL database connection arguments: %w", err)
	}

	err = db.Ping()
	if err != nil {
		return MysqlStorage{}, fmt.Errorf("cannot open MySQL database connection: %w", err)
	}

	mysqlStorage := MysqlStorage{
		db: db,
	}
	log.Printf("MySQL database connection created: %v", mysqlStorage)

	return mysqlStorage, nil
}
