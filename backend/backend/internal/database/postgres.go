package database

import (
	"database/sql"

	licensesqueries "github.com/fossology/licenseDb/backend/internal/database/queries"
)

func NewQueries(db *sql.DB) *licensesqueries.Queries {
	return licensesqueries.New(db)
}
