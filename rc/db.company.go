package rc

import "github.com/gin-gonic/gin"

func fetchAllCompanies(ctx *gin.Context, rid string, companies *[]CompanyRecruitmentCycle) error {
	tx := db.WithContext(ctx).Where("recruitment_cycle_id = ?", rid).Find(companies)
	return tx.Error
}

func fetchCompany(ctx *gin.Context, rid string, cid string, company *CompanyRecruitmentCycle) error {
	tx := db.WithContext(ctx).Where("recruitment_cycle_id = ? AND company_id = ?", rid, cid).First(company)
	return tx.Error
}

func FetchCompanyByID(ctx *gin.Context, cid uint, company *CompanyRecruitmentCycle) error {
	tx := db.WithContext(ctx).Where("id = ?", cid).First(company)
	return tx.Error
}

func createCompany(ctx *gin.Context, company *CompanyRecruitmentCycle) error {
	tx := db.WithContext(ctx).Create(company)
	return tx.Error
}

func FetchCompanyRCID(ctx *gin.Context, rid uint, companyid uint) (uint, error) {
	var company CompanyRecruitmentCycle
	tx := db.WithContext(ctx).Where("recruitment_cycle_id = ? AND company_id = ?", rid, companyid).First(&company)
	return company.ID, tx.Error
}

func getRegisteredCompanyCount(ctx *gin.Context, rid uint) (int, error) {
	var count int64
	tx := db.WithContext(ctx).Model(&CompanyRecruitmentCycle{}).Where("recruitment_cycle_id = ?", rid).Count(&count)
	return int(count), tx.Error
}