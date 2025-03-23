package system

import (
	licensesdao "github.com/fossology/licenseDb/backend/internal/dao/licenses"
	"github.com/fossology/licenseDb/backend/internal/database"
	licenseshandler "github.com/fossology/licenseDb/backend/internal/handler/licenses"
	licensesrepo "github.com/fossology/licenseDb/backend/internal/repo/licenses"
	"github.com/fossology/licenseDb/backend/internal/router"
	"github.com/fossology/licenseDb/backend/pkg/postgres"
)

func InitDependencies() {
	db := postgres.NewPostgresDB() // Connect to DB

	queryExecutor := database.NewQueries(db)

	licensesDAO := licensesdao.NewLicensesDAO(db, queryExecutor)
	licensesRepo := licensesrepo.NewLicensesRepo(licensesDAO)
	licensesHandler := licenseshandler.NewLicensesHandler(licensesRepo)

	router.InitRoutes(licensesHandler) // Inject handler into router
}
