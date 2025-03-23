package licensesrepo

import (
	licensesdao "github.com/fossology/licenseDb/backend/internal/dao/licenses"
	licensesModel "github.com/fossology/licenseDb/backend/internal/models/licenses"
)

type LicensesRepo struct {
	dao *licensesdao.LicensesDAO
}

func NewLicensesRepo(dao *licensesdao.LicensesDAO) *LicensesRepo {
	return &LicensesRepo{dao: dao}
}

func (repo *LicensesRepo) GetLicense(id int64) (*licensesModel.License, error) {
	return repo.dao.GetLicenseByIDRaw(id)
}
