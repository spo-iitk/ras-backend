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

func createCompany(ctx *gin.Context, company *Company) error {
	tx := db.WithContext(ctx).Create(company)
	return tx.Error
}

func createCompanies(ctx *gin.Context, company *[]Company) error {
	tx := db.WithContext(ctx).Create(company)
	return tx.Error
}

func deleteCompany(ctx *gin.Context, id uint) error {
	tx := db.WithContext(ctx).Where("id = ?", id).Delete(&Company{})
	return tx.Error
}

func FetchCompanyIDByEmail(ctx *gin.Context, email string) (uint, error) {
	var company Company
	tx := db.WithContext(ctx).Where("email = ?", email).First(&company)
	return company.ID, tx.Error
}
