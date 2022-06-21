package rc

import (
	"github.com/gin-gonic/gin"
)

func fetchAllRCs(ctx *gin.Context, rc *[]RecruitmentCycle) error {
	tx := db.WithContext(ctx).Find(rc)
	return tx.Error
}

func fetchRCsByStudent(ctx *gin.Context, email string, rcs *[]RecruitmentCycle) error {
	tx := db.WithContext(ctx).
		Joins("JOIN student_recruitment_cycles ON student_recruitment_cycles.recruitment_cycle_id = recruitment_cycles.id AND student_recruitment_cycles.email = ?", email).Find(&rcs)
	return tx.Error
}

func fetchRCsByCompanyID(ctx *gin.Context, cid uint, rcs *[]RecruitmentCycle) error {
	tx := db.WithContext(ctx).
		Joins("JOIN ON company_recruitment_cycles ON company_recruitment_cycles.recruitment_cycle_id = recruitment_cycles.id AND company_recruitment_cycles.company_id = ?", cid).Find(&rcs)
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
