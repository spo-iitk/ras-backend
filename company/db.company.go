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

func getLimitedCompanies(ctx *gin.Context, companies *[]Company, lastFetchedId uint, pageSize int) error {
	tx := db.WithContext(ctx).Order("id asc").Where("id >= ?", lastFetchedId).Limit(pageSize).Find(companies)
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

func GetCompanyName(ctx *gin.Context, id uint) (string, error) {
	var c Company
	err := getCompany(ctx, &c, id)
	if err != nil {
		return "", err
	}
	return c.Name, nil
}
