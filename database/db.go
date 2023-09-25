package database

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
)

func InitDatabase() (*sql.DB, error) {
	connStr := os.Getenv("DATABASE_URL") 
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	return db, nil
}
