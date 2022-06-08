package company

import (
	"github.com/gin-gonic/gin"
)

func getAllCompanies(ctx *gin.Context, companies *[]Company) error {
	tx := db.WithContext(ctx).Find(companies)
	return tx.Error
}

func getCompany(ctx *gin.Context, company *Company, id uint) error {
	tx := db.WithContext(ctx).Where("id = ?", id).First(company)
	return tx.Error
}

func updateCompany(ctx *gin.Context, company *Company) (bool, error) {
	tx := db.WithContext(ctx).Where("id = ?", company.ID).Updates(company)
	return tx.RowsAffected > 0, tx.Error
}

func newCompany(ctx *gin.Context, company *Company) error {
	tx := db.WithContext(ctx).Create(company)
	return tx.Error
}

func deleteCompany(ctx *gin.Context, id uint) error {
	tx := db.WithContext(ctx).Where("id = ?", id).Delete(&Company{})
	return tx.Error
}

func getAllHR(ctx *gin.Context, HRs *[]CompanyHR, cid uint) error {
	tx := db.WithContext(ctx).Where("company_id = ?", cid).Find(HRs)
	return tx.Error
}

func addHR(ctx *gin.Context, HR *CompanyHR) error {
	tx := db.WithContext(ctx).Create(HR)
	return tx.Error
}

func getHR(ctx *gin.Context, HR *CompanyHR, id uint) error {
	tx := db.WithContext(ctx).Where("id = ?", id).First(HR)
	return tx.Error
}

func updateHR(ctx *gin.Context, HR *CompanyHR) (bool, error) {
	tx := db.WithContext(ctx).Where("id = ?", HR.ID).Updates(HR)
	return tx.RowsAffected > 0, tx.Error
}

func deleteHR(ctx *gin.Context, id uint) error {
	tx := db.WithContext(ctx).Delete(&CompanyHR{}, "id = ?", id)
	return tx.Error
}
