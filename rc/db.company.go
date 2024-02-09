package rc

import "github.com/gin-gonic/gin"
import "fmt"

func fetchAllCompanies(ctx *gin.Context, rid string, companies *[]CompanyRecruitmentCycle) error {
	tx := db.WithContext(ctx).Where("recruitment_cycle_id = ?", rid).Find(companies)
	return tx.Error
}

func fetchCompanyByRCIDAndCID(ctx *gin.Context, rid uint, cid uint, company *CompanyRecruitmentCycle) error {
	tx := db.WithContext(ctx).Where("recruitment_cycle_id = ? AND company_id = ?", rid, cid).First(company)
	return tx.Error
}

func FetchCompany(ctx *gin.Context, cid uint, company *CompanyRecruitmentCycle) error {
	tx := db.WithContext(ctx).Where("id = ?", cid).First(company)
	return tx.Error
}

func upsertCompany(ctx *gin.Context, company *CompanyRecruitmentCycle) error {
	tx := db.WithContext(ctx).
		Where("company_id = ? AND recruitment_cycle_id = ?", company.CompanyID, company.RecruitmentCycleID).
		Updates(company)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		tx = db.WithContext(ctx).
			Where("company_id = ? AND recruitment_cycle_id = ?", company.CompanyID, company.RecruitmentCycleID).
			FirstOrCreate(company)
	}
	return tx.Error
}

func editCompany(ctx *gin.Context, company *CompanyRecruitmentCycle) error {
	tx := db.WithContext(ctx).Where("id = ?", company.ID).Updates(company)
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

func deleteRCCompany(ctx *gin.Context, cid uint) error {
	tx := db.WithContext(ctx).Where("id = ?", cid).Delete(&CompanyRecruitmentCycle{})
	return tx.Error
}

func FetchCompanyHistory(ctx *gin.Context, companyID uint, companyHistory *[]CompanyHistory) error {
    tx := db.WithContext(ctx).
        Table("company_recruitment_cycles").
        Select("company_recruitment_cycles.id, company_recruitment_cycles.recruitment_cycle_id, recruitment_cycles.type, recruitment_cycles.phase, company_recruitment_cycles.comments").
        Joins("JOIN recruitment_cycles ON company_recruitment_cycles.recruitment_cycle_id = recruitment_cycles.id").
        Where("company_recruitment_cycles.company_id = ?", companyID).
        Find(companyHistory)

    if tx.Error != nil {
        fmt.Println("Error fetching company history:", tx.Error)
        return tx.Error
    }

    return nil
}