package rc

import "github.com/gin-gonic/gin"

func addStudentResume(ctx *gin.Context, resume string, sid uint, rid uint) error {
	tx := db.WithContext(ctx).Model(&StudentRecruitmentCycleResume{}).Create(&StudentRecruitmentCycleResume{
		StudentRecruitmentCycleID: sid,
		Resume:                    resume,
		RecruitmentCycleID:        rid,
	})
	return tx.Error
}

func fetchStudentResume(ctx *gin.Context, sid uint) ([]string, error) {
	var resumes []string
	tx := db.WithContext(ctx).Model(&StudentRecruitmentCycleResume{}).Where("student_recruitment_cycle_id = ?", sid).Pluck("resume", &resumes)
	return resumes, tx.Error
}

func fetchAllResumes(ctx *gin.Context, rid uint, resumes *[]StudentRecruitmentCycleResume) error {
	tx := db.WithContext(ctx).Where("recruitment_cycle_id = ?", rid).Find(resumes)
	return tx.Error
}

func fetchResume(ctx *gin.Context, rsid uint) (string, error) {
	var resume string
	tx := db.WithContext(ctx).Model(&StudentRecruitmentCycleResume{}).Where("id = ?", rsid).Pluck("resume", &resume)
	return resume, tx.Error
}

func updateResumeVerify(ctx *gin.Context, rsid uint, verified bool) (bool, error) {
	tx := db.WithContext(ctx).Model(&StudentRecruitmentCycleResume{}).Where("id = ?", rsid).Update("verified", verified)
	return tx.RowsAffected == 1, tx.Error
}
