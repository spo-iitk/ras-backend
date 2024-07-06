package application

import (
	"github.com/gin-gonic/gin"
)

func createPVF(ctx *gin.Context, pvf *PVF) error {
	tx := db.WithContext(ctx).Create(pvf)
	return tx.Error
}

func fetchPVF(ctx *gin.Context, pid uint, jp *PVF) error {
	tx := db.WithContext(ctx).Where("id = ?", pid).First(jp)
	return tx.Error
}

func updatePVF(ctx *gin.Context, jp *PVF) error {
	tx := db.WithContext(ctx).Where("id = ?", jp.ID).Updates(jp)
	return tx.Error
}

func fetchAllPvfForStudent(ctx *gin.Context, sid uint, rid uint, jps *[]PVF) error {
	tx := db.WithContext(ctx).
		Where("student_recruitment_cycle_id = ? AND recruitment_cycle_id = ?", sid, rid).
		Order("id DESC").
		Find(jps)
	return tx.Error
}
func fetchPvfForStudent(ctx *gin.Context, sid uint, rid uint, pid uint, jps *PVF) error {
	tx := db.WithContext(ctx).
		Where("student_recruitment_cycle_id = ? AND recruitment_cycle_id = ? AND id = ?", sid, rid, pid).
		Select(
			"id",
			"company_university_name",
			"role",
			"duration",
			"mentor_name",
			"mentor_designation",
			"mentor_email",
			"is_verified",
			"filename_mentor",
			"filename_student",
			"remarks",
		).
		Find(jps)
	return tx.Error
}
func fetchPvfForAdmin(ctx *gin.Context, rid uint, pid uint, jps *PVF) error {
	tx := db.WithContext(ctx).
		Where("recruitment_cycle_id = ? AND id = ?", rid, pid).
		Order("id DESC").
		Find(jps)
	return tx.Error
}

func fetchPvfForVerification(ctx *gin.Context, id uint, rid uint, jps *PVF) error {
	tx := db.WithContext(ctx).
		Where("id = ? AND recruitment_cycle_id = ?", id, rid).
		Select(
			"id",
			"company_university_name",
			"role",
			"duration",
			"mentor_name",
			"mentor_designation",
			"mentor_email",
			"is_verified",
			"is_approved",
			"filename_student",
			"filename_mentor",
			"remarks",
		).
		Find(jps)
	return tx.Error
}

func fetchAllPvfForAdmin(ctx *gin.Context, rid uint, jps *[]PVF) error {
	tx := db.WithContext(ctx).
		Where("recruitment_cycle_id = ?", rid).
		Order("id DESC").
		Find(jps)
	return tx.Error
}

func deletePVF(ctx *gin.Context, pid uint) error {
	tx := db.WithContext(ctx).Where("id = ?", pid).Delete(&PVF{})
	return tx.Error
}

// func fetchAllStudentPvfForAdmin(ctx *gin.Context)

func updatePVFForStudent(ctx *gin.Context, sid uint, jp *PVF) (bool, error) {
	tx := db.WithContext(ctx).Where("id = ? AND student_recruitment_cycle_id = ?", jp.ID, sid).Updates(jp)
	return tx.RowsAffected == 1, tx.Error
}
