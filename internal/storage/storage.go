package storage

import (
	"database/sql"
)

type SQLiteTransport struct {
	db   *sql.DB
	path string
}

func NewSQLiteTransport(scope, basePath string) (*SQLiteTransport, error) {
	// Implement SQLite connection logic
	// ...
}

func (s *SQLiteTransport) Close() error {
	return s.db.Close()
}

func (s *SQLiteTransport) GetAllObjects() ([]StorageObject, error) {
	// Implement logic to retrieve all objects from SQLite
	// ...
}

type StorageObject struct {
	ID   string
	Data string
}
