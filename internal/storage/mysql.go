package storage

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"gopi/internal/lib/random"
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

func (s *Storage) SaveGif(uuid string, path string) (int64, string, error) {
	const op = "storage.mysql.SaveGif"

	alias := random.NewRandomString()

	stmt, err := s.Db.Prepare("INSERT INTO gifs (uuid, path, alias) VALUES(?, ?, ?)")
	if err != nil {
		return 0, "", fmt.Errorf("%s: %w", op, err)
	}
	defer stmt.Close()

	res, err := stmt.Exec(uuid, path, alias)
	if err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) {
			return 0, "", fmt.Errorf("UUID or alias already exists: %w", err)
		}
		return 0, "", fmt.Errorf("failed to execute statement: %w", err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, "", fmt.Errorf("failed to get last insert id: %w", err)
	}

	return id, alias, nil
}

func (s *Storage) GetGifByAliasOrUUID(id string) (string, error) {
	var path string
	query := "SELECT path FROM gifs WHERE uuid = ? OR alias = ? LIMIT 1"

	err := s.Db.QueryRow(query, id, id).Scan(&path)
	if errors.Is(err, sql.ErrNoRows) {
		return "", fmt.Errorf("gif not found for identifier: %s", id)
	} else if err != nil {
		return "", fmt.Errorf("failed to query database: %w", err)
	}

	return path, nil
}

func (s *Storage) GetAllAliases() ([]string, error) {
	const op = "storage.mysql.GetAllAliases"

	query := "SELECT alias FROM gifs"

	rows, err := s.Db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("%s: failed to query database: %w", op, err)
	}
	defer rows.Close()

	var aliases []string

	for rows.Next() {
		var alias string

		if err := rows.Scan(&alias); err != nil {
			return nil, fmt.Errorf("%s: failed to scan row: %w", op, err)
		}

		aliases = append(aliases, alias)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("%s: row iteration error: %w", op, err)
	}

	if len(aliases) == 0 {
		return nil, fmt.Errorf("%s: no aliases found", op)
	}

	return aliases, nil
}
