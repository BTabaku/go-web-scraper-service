// pkg/database/database.go
package database

import (
	"database/sql"
	"log"
	// PostgreSQL driver
)

type Database struct {
	Conn *sql.DB
}

func NewDatabase(connStr string) (*Database, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	// Test connection
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	log.Println("Database connected successfully")
	return &Database{Conn: db}, nil
}

func (db *Database) SavePage(url, content string) error {
	_, err := db.Conn.Exec("INSERT INTO pages (url, content) VALUES ($1, $2)", url, content)
	return err
}
