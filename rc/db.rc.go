package rc

import (
	"github.com/gin-gonic/gin"
)

func fetchAllRCs(ctx *gin.Context, rc *[]RecruitmentCycle) error {
	tx := db.WithContext(ctx).Find(rc)
	return tx.Error
}

func fetchAllActiveRCs(ctx *gin.Context, rc *[]RecruitmentCycle) error {
	tx := db.WithContext(ctx).Where("is_active = ?", true).Find(rc)
	return tx.Error
}

func fetchRCsByStudent(ctx *gin.Context, email string, rcs *[]RecruitmentCycle) error {
	tx := db.WithContext(ctx).
		Joins("JOIN student_recruitment_cycles ON student_recruitment_cycles.recruitment_cycle_id = recruitment_cycles.id").
		Where("student_recruitment_cycles.deleted_at IS NULL AND student_recruitment_cycles.email = ?", email).
		Find(&rcs)
	return tx.Error
}

func fetchRCsByCompanyID(ctx *gin.Context, cid uint, rcs *[]RecruitmentCycle) error {
	tx := db.WithContext(ctx).
		Joins("JOIN company_recruitment_cycles ON company_recruitment_cycles.recruitment_cycle_id = recruitment_cycles.id").
		Where("company_recruitment_cycles.deleted_at IS NULL AND company_recruitment_cycles.company_id = ? AND recruitment_cycles.deleted_at IS NULL AND recruitment_cycles.is_active = ?", cid, true).Find(&rcs)
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

func IsRCActive(ctx *gin.Context, rid uint) bool {
	tx := db.WithContext(ctx).Where("id = ? AND is_active = ?", rid, true).First(&RecruitmentCycle{})
	return tx.Error == nil
}

func updateRC(ctx *gin.Context, id uint, inactive bool, countcap uint) (bool, error) {
	tx := db.WithContext(ctx).Model(&RecruitmentCycle{}).Where("id = ?", id).
		Update("is_active", !inactive).Updates(&RecruitmentCycle{ApplicationCountCap: countcap})
	return tx.RowsAffected == 1, tx.Error
}
