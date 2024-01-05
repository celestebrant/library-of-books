package storage

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/celestebrant/docker-sql-demo/book"
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

// CreateBook creates a record of a book in the Mysql storage.
func (s *MysqlStorage) CreateBook(ctx context.Context, b *book.Book) (*book.Book, error) {
	query := "INSERT INTO `books` (`creation_time`, `title`, `author`) VALUES (?, ?, ?);"

	if b.CreationTime.IsZero() {
		b.CreationTime = time.Now()
	}

	result, err := s.db.ExecContext(ctx, query, b.CreationTime, b.Title, b.Author)
	if err != nil {
		return &book.Book{}, fmt.Errorf("error during insert: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return &book.Book{}, fmt.Errorf("cannot get last insert id: %w", err)
	}

	b.ID = id

	return b, nil
}
