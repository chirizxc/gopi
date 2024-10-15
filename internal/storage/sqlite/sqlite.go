package sqlite

import (
	"database/sql"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"time"

	_uuid "github.com/google/uuid"
	"github.com/mattn/go-sqlite3"
	"gopi/internal/storage"
)

type Storage struct {
	db *sql.DB
}

func New(storagePath string) (*Storage, error) {
	const op = "storage.sqlite.New"

	db, err := sql.Open("sqlite3", storagePath)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	stmt, err := db.Prepare(`CREATE TABLE IF NOT EXISTS url(
		id INTEGER PRIMARY KEY,
		uuid TEXT NOT NULL UNIQUE,
		alias TEXT NOT NULL UNIQUE,
		url TEXT NOT NULL,
		gif_path TEXT NOT NULL);
	CREATE INDEX IF NOT EXISTS idx_alias ON url(alias);`)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	_, err = stmt.Exec()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &Storage{db: db}, nil
}

func (s *Storage) SaveURL(urlToSave string, gifFilePath string) (int64, error) {
	const op = "storage.sqlite.SaveURL"

	uuid := _uuid.New().String()
	alias := generateAlias()

	if _, err := os.Stat(gifFilePath); os.IsNotExist(err) {
		return 0, fmt.Errorf("%s: gif file does not exist: %w", op, err)
	}

	stmt, err := s.db.Prepare("INSERT INTO url(uuid, url, alias, gif_path) VALUES(?, ?, ?, ?)")
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	res, err := stmt.Exec(uuid, urlToSave, alias, gifFilePath)
	if err != nil {
		var sqliteErr sqlite3.Error
		if errors.As(err, &sqliteErr) && errors.Is(sqliteErr.ExtendedCode, sqlite3.ErrConstraintUnique) {
			return 0, fmt.Errorf("%s: %w", op, storage.URLExists)
		}
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("%s: failed to get last insert id: %w", op, err)
	}

	return id, nil
}

func (s *Storage) GetURL(alias string) (string, string, error) {
	const op = "storage.sqlite.GetURL"

	stmt, err := s.db.Prepare("SELECT url, gif_path FROM url WHERE alias = ?")
	if err != nil {
		return "", "", fmt.Errorf("%s: prepare statement: %w", op, err)
	}

	var resURL, gifPath string

	err = stmt.QueryRow(alias).Scan(&resURL, &gifPath)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", "", storage.URLNotFound
		}
		return "", "", fmt.Errorf("%s: execute statement: %w", op, err)
	}

	return resURL, gifPath, nil
}

func generateAlias() string {
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	alias := make([]byte, 8)
	for i := range alias {
		alias[i] = charset[r.Intn(len(charset))]
	}
	return string(alias)
}

// TODO: need make DeleteURL func to remove the URL
// https://t.me/c/2420815282/2
