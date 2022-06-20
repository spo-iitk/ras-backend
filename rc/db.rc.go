package rc

import (
	"github.com/gin-gonic/gin"
)

func fetchAllRCs(ctx *gin.Context, rc *[]RecruitmentCycle) error {
	tx := db.WithContext(ctx).Find(rc)
	return tx.Error
}

func fetchRCsByStudent(ctx *gin.Context, email string, rcs *[]RecruitmentCycle) error {
	tx := db.WithContext(ctx).Joins("StudentRecruitmentCycle", db.Where(&StudentRecruitmentCycle{Email: email})).Find(&rcs)
	return tx.Error
}

func fetchRCsByCompanyID(ctx *gin.Context, cid uint, rcs *[]RecruitmentCycle) error {
	tx := db.WithContext(ctx).Joins("company_recruitment_cycle", db.Where(&CompanyRecruitmentCycle{CompanyID: cid})).Find(&rcs)
	return tx.Error
}

func createRC(ctx *gin.Context, rc *RecruitmentCycle) error {
	tx := db.WithContext(ctx).Create(rc)
	return tx.Error
}

func fetchRC(ctx *gin.Context, rid string, rc *RecruitmentCycle) error {
	tx := db.WithContext(ctx).Where("id = ?", rid).First(rc)
	return tx.Error
}
