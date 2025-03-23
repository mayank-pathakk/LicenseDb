package postgres

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func NewClient(connStr string) *sql.DB {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("failed to connect to postgres: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("failed to ping postgres: %v", err)
	}

	return db
}

func NewPostgresDB() *sql.DB {
	connStr := "postgres://fossy:fossy@localhost:5432/example?sslmode=disable"
	return NewClient(connStr)
}
