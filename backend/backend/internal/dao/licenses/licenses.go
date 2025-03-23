package licenses

import (
	"context"
	"database/sql"

	licensesqueries "github.com/fossology/licenseDb/backend/internal/database/queries"
	licensesModel "github.com/fossology/licenseDb/backend/internal/models/licenses"
)

type LicensesDAO struct {
	db      *sql.DB
	queries *licensesqueries.Queries
}

func NewLicensesDAO(db *sql.DB, queries *licensesqueries.Queries) *LicensesDAO {
	return &LicensesDAO{db: db, queries: queries}
}

// Approach 1: Raw query
func (dao *LicensesDAO) GetLicenseByIDRaw(id int64) (*licensesModel.License, error) {
	row := dao.db.QueryRow("SELECT id, name FROM licenses WHERE id = $1", id)
	var license licensesModel.License
	err := row.Scan(&license.ID, &license.Name)
	return &license, err
}

// Approach 2: Using sqlc-generated query
func (dao *LicensesDAO) GetLicenseByID(id int64) (*licensesModel.License, error) {
	license, err := dao.queries.GetLicenseByID(context.Background(), int32(id))
	if err != nil {
		return nil, err
	}
	return &licensesModel.License{
		ID:   int64(license.ID),
		Name: license.Name,
	}, nil
}
