package router

import (
	"net/http"

	licenseshandler "github.com/fossology/licenseDb/backend/internal/handler/licenses"
)

func InitRoutes(handler *licenseshandler.LicensesHandler) {
	http.HandleFunc("/license", handler.GetLicense)
}
