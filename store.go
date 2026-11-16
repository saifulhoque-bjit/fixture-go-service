package main

import (
	"database/sql"
	"errors"
	"strings"

	_ "modernc.org/sqlite"
)

// User is one entry of the directory.
type User struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// ErrNotFound means no user has the name.
var ErrNotFound = errors.New("user not found")

// Store keeps users in SQLite.
type Store struct {
	db *sql.DB
}

// OpenStore opens the database and creates the table.
func OpenStore(dsn string) (*Store, error) {
	db, err := sql.Open("sqlite", dsn)
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(1)
	if _, err := db.Exec(`CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL UNIQUE,
		email TEXT NOT NULL
	)`); err != nil {
		db.Close()
		return nil, err
	}
	return &Store{db: db}, nil
}

// Close closes the database.
func (s *Store) Close() error { return s.db.Close() }

// NormalizeName is the form a name is stored and looked up in.
func NormalizeName(name string) string {
	return strings.TrimSpace(name)
}

// AddUser stores a user.
func (s *Store) AddUser(name, email string) (User, error) {
	res, err := s.db.Exec("INSERT INTO users (name, email) VALUES (?, ?)", NormalizeName(name), email)
	if err != nil {
		return User{}, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return User{}, err
	}
	return User{ID: id, Name: NormalizeName(name), Email: email}, nil
}
