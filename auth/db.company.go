package auth

import "github.com/gin-gonic/gin"

func createCompany(ctx *gin.Context, company *CompanySignUpRequest) (uint, error) {
	tx := db.WithContext(ctx).Create(company)
	return company.ID, tx.Error
}
