package auth

import "github.com/gin-gonic/gin"

func createCompany(ctx *gin.Context, company *CompanySignUpRequest) (uint, error) {
	tx := db.WithContext(ctx).Create(company)
	return company.ID, tx.Error
}

func getAllCompaniesAdded(ctx *gin.Context) ([]string, error) {
	var companies []string
	tx := db.WithContext(ctx).Model(&CompanySignUpRequest{}).Order("created_at DESC").
		Limit(50).Pluck("company_name", &companies)
	return companies, tx.Error
}
