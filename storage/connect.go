package storage

import (
	"database/sql"
	"fmt"

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
		return MysqlStorage{}, fmt.Errorf("cannot validate open SQL connection arguments: %w", err)
	}

	err = db.Ping()
	if err != nil {
		return MysqlStorage{}, fmt.Errorf("cannot open SQL connection: %w", err)
	}

	return MysqlStorage{
		db: db,
	}, nil
}
