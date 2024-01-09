package storage

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/celestebrant/docker-sql-demo/book"
	_ "github.com/go-sql-driver/mysql" // blank import runs init
	"github.com/oklog/ulid/v2"
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

// CreateBook creates a record of a book in the Mysql storage. If book creation time has a zero value,
// then it will be set to now. If the book ID is empty, then it will take on a generated ULID string.
func (s *MysqlStorage) CreateBook(ctx context.Context, b *book.Book) (*book.Book, error) {
	query := "INSERT INTO `books` (`id`, `creation_time`, `title`, `author`) VALUES (?, ?, ?, ?);"

	if b.CreationTime.IsZero() {
		b.CreationTime = time.Now()
	}
	if b.ID == "" {
		b.ID = ulid.Make().String()
	}

	err := b.Validate()
	if err != nil {
		return nil, fmt.Errorf("cannot insert book: %w", err)
	}

	_, err = s.db.ExecContext(ctx, query, b.ID, b.CreationTime, b.Title, b.Author)
	if err != nil {
		return &book.Book{}, fmt.Errorf("error during insert: %w", err)
	}

	return b, nil
}
