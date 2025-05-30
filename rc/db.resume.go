package rc

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

func addStudentResume(ctx *gin.Context, resume string, sid uint, rid uint, resumeType ResumeType, resumeTag string) error {
	tx := db.WithContext(ctx).Model(&StudentRecruitmentCycleResume{}).Create(&StudentRecruitmentCycleResume{
		StudentRecruitmentCycleID: sid,
		Resume:                    resume,
		RecruitmentCycleID:        rid,
		ResumeType:                resumeType,
		Tag:                       resumeTag, // ← Store the tag
	})
	return tx.Error
}

func fetchStudentResume(ctx *gin.Context, sid uint, resumes *[]StudentRecruitmentCycleResume) error {
	tx := db.WithContext(ctx).Model(&StudentRecruitmentCycleResume{}).Where("student_recruitment_cycle_id = ?", sid).Find(resumes)
	return tx.Error
}
func fetchAllResumes(ctx *gin.Context, rid uint, resumes *[]AllResumeResponse) error {
	tx := db.WithContext(ctx).Model(&StudentRecruitmentCycleResume{}).
		Joins("JOIN student_recruitment_cycles ON student_recruitment_cycles.id = student_recruitment_cycle_resumes.student_recruitment_cycle_id AND student_recruitment_cycle_resumes.recruitment_cycle_id = ?", rid).
		Select("student_recruitment_cycles.name as name, student_recruitment_cycles.email as email, student_recruitment_cycles.id as sid, student_recruitment_cycle_resumes.id as rsid, student_recruitment_cycles.roll_no as roll_no, student_recruitment_cycle_resumes.resume as resume, student_recruitment_cycle_resumes.verified as verified, student_recruitment_cycle_resumes.action_taken_by as action_taken_by, student_recruitment_cycle_resumes.resume_type as resume_type"). // Include resume_type in the response
		Scan(resumes)
	return tx.Error
}

func FetchResume(ctx *gin.Context, rsid uint, sid uint) (string, error) {
	var resume string
	tx := db.WithContext(ctx).Model(&StudentRecruitmentCycleResume{}).
		Where("id = ? AND student_recruitment_cycle_id = ? AND verified = ?", rsid, sid, true).
		Pluck("resume", &resume)
	return resume, tx.Error
}

func FetchFirstResume(ctx *gin.Context, sid uint) (uint, string, error) {
	var resume StudentRecruitmentCycleResume
	tx := db.WithContext(ctx).Model(&StudentRecruitmentCycleResume{}).
		Where("student_recruitment_cycle_id = ? AND verified = ?", sid, true).First(&resume)
	return resume.ID, resume.Resume, tx.Error
}

func updateResumeVerify(ctx *gin.Context, rsid uint, verified bool, user string) (bool, uint, error) {
	var resume StudentRecruitmentCycleResume
	tx := db.WithContext(ctx).Model(&resume).Clauses(clause.Returning{}).
		Where("id = ?", rsid).Updates(map[string]interface{}{"verified": verified, "action_taken_by": user})
	return tx.RowsAffected == 1, resume.StudentRecruitmentCycleID, tx.Error
}
