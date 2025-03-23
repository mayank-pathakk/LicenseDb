package licenseshandler

import (
	"encoding/json"
	"net/http"
	"strconv"

	licensesRepo "github.com/fossology/licenseDb/backend/internal/repo/licenses"
)

type LicensesHandler struct {
	repo *licensesRepo.LicensesRepo
}

func NewLicensesHandler(repo *licensesRepo.LicensesRepo) *LicensesHandler {
	return &LicensesHandler{repo: repo}
}

func (h *LicensesHandler) GetLicense(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	license, err := h.repo.GetLicense(id)
	if err != nil {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(license)
}
