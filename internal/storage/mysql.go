package storage

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"time"
)

type DB struct {
	*sql.DB
}

type Storage struct {
	Db *DB
}

func NewDB(dsn string) (*DB, error) {

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return &DB{db}, nil
}

func (s *Storage) SaveGif(uuid string, path string) (int64, error) {
	const op = "storage.mysql.SaveGif"

	stmt, err := s.Db.Prepare("INSERT INTO gifs (uuid, path) VALUES(?, ?)")
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}
	defer stmt.Close()

	res, err := stmt.Exec(uuid, path)
	if err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) {
			return 0, fmt.Errorf("UUID already exists: %w", err)
		}
		return 0, fmt.Errorf("failed to execute statement: %w", err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("failed to get last insert id: %w", err)
	}

	return id, nil
}
