package application

import "github.com/gin-gonic/gin"

func createPVF(ctx *gin.Context, pvf *PVF) error {
	tx := db.WithContext(ctx).Create(pvf)
	return tx.Error
}

func fetchPvfForStudent(ctx *gin.Context, sid uint, rid uint, jps *[]PVF) error {
	tx := db.WithContext(ctx).
		// Where("student_recruitment_cycle_id = ? AND recruitment_cycle_id = ?", sid, rid).
		Select(
			"id",
			"company_university_name",
			"role",
			"duration",
			"description",
			"mentor_name",
			"mentor_designation",
			"mentor_email",
			"is_verified",
		).
		Order("id ASC").
		Find(jps)
	return tx.Error
}
