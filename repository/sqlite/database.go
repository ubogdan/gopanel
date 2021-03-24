// +build sqlite_foreign_keys

package sqlite

import (
	"database/sql"
)

// Database return new database.
func Database(sqlite3FileName string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", sqlite3FileName+"?cache=shared&mode=ro")
	if err != nil {
		return nil, err
	}

	_, err = db.Exec("PRAGMA foreign_keys=ON;")
	if err != nil {
		return nil, err
	}

	return db, nil
}

// NewMemoryDatabase return in memory database.
func NewMemoryDatabase() (*sql.DB, error) {
	return Database(":memory:")
}
