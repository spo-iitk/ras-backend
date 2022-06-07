package rc

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func fetchAllRCs(ctx *gin.Context, rc *[]RecruitmentCycle) error {
	tx := db.WithContext(ctx).Find(&rc)
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

func fetchAllNotices(ctx *gin.Context, rid string, notices *[]Notice) error {
	tx := db.WithContext(ctx).Where("recruitment_cycle_id = ?", rid).Find(&notices)
	return tx.Error
}

func createNotice(ctx *gin.Context, notice *Notice) error {
	tx := db.WithContext(ctx).Create(notice)
	return tx.Error
}

func removeNotice(ctx *gin.Context, nid string) error {
	tx := db.WithContext(ctx).Where("id = ?", nid).Delete(Notice{})
	if tx.RowsAffected == 0 {
		return errors.New("no notice found")
	}
	return tx.Error
}

func updateNotice(ctx *gin.Context, notice *Notice) error {
	tx := db.WithContext(ctx).Save(notice)
	return tx.Error
}

func fetchNotice(ctx *gin.Context, nid string, notice *Notice) error {
	tx := db.WithContext(ctx).Where("id = ?", nid).First(notice)
	return tx.Error
}

func fetchAllCompanies(ctx *gin.Context, rid string, companies *[]CompanyRecruitmentCycle) error {
	tx := db.WithContext(ctx).Where("recruitment_cycle_id = ?", rid).Find(&companies)
	return tx.Error
}

func fetchCompany(ctx *gin.Context, rid string, cid string, company *CompanyRecruitmentCycle) error {
	tx := db.WithContext(ctx).Where("recruitment_cycle_id = ? AND company_id = ?", rid, cid).First(&company)
	return tx.Error
}

func createCompany(ctx *gin.Context, company *CompanyRecruitmentCycle) error {
	tx := db.WithContext(ctx).Create(&company)
	return tx.Error
}

func fetchAllStudents(ctx *gin.Context, rid string, students *[]StudentRecruitmentCycle) error {
	tx := db.WithContext(ctx).Where("recruitment_cycle_id = ?", rid).Find(&students)
	return tx.Error
}

func fetchStudent(ctx *gin.Context, rid string, sid string, student *StudentRecruitmentCycle) error {
	tx := db.WithContext(ctx).Where("recruitment_cycle_id = ? AND student_id = ?", rid, sid).First(&student)
	return tx.Error
}

func updateStudent(ctx *gin.Context, student *StudentRecruitmentCycle) error {
	tx := db.WithContext(ctx).Save(&student)
	return tx.Error
}

func createStudents(ctx *gin.Context, students *[]StudentRecruitmentCycle) error {
	tx := db.WithContext(ctx).Create(&students)
	return tx.Error
}

func fetchStudentQuestions(ctx *gin.Context, rid string, sid string, questions *[]RecruitmentCycleQuestion) error {
	tx := db.WithContext(ctx).Where("recruitment_cycle_id = ? AND student_recruitment_cycle_id = ?", rid, sid).Find(&questions)
	return tx.Error
}

func createStudentQuestion(ctx *gin.Context, question *RecruitmentCycleQuestion) error {
	tx := db.WithContext(ctx).Create(&question)
	return tx.Error
}

func updateStudentQuestion(ctx *gin.Context, question *RecruitmentCycleQuestion) error {
	tx := db.WithContext(ctx).Save(&question)
	return tx.Error
}

func deleteStudentQuestion(ctx *gin.Context, qid string) error {
	tx := db.WithContext(ctx).Where("id = ?", qid).Delete(&RecruitmentCycleQuestion{})
	return tx.Error
}

func fetchStudentAnswer(ctx *gin.Context, qid string, sid string, answer *RecruitmentCycleQuestionsAnswer) error {
	tx := db.WithContext(ctx).Where("recruitment_cycle_question_id = ? AND student_id = ?", qid, sid).Find(&answer)
	return tx.Error
}

func createStudentAnswer(ctx *gin.Context, answer *RecruitmentCycleQuestionsAnswer) error {
	tx := db.WithContext(ctx).Create(&answer)
	return tx.Error
}

func updateStudentAnswer(ctx *gin.Context, answer *RecruitmentCycleQuestionsAnswer) error {
	tx := db.WithContext(ctx).Save(&answer)
	return tx.Error
}

func deleteStudentAnswer(ctx *gin.Context, qid string, sid string) error {
	tx := db.WithContext(ctx).Where("recruitment_cycle_question_id = ? AND student_id = ?", qid, sid).Delete(&RecruitmentCycleQuestionsAnswer{})
	return tx.Error
}

func updateStudentType(ctx *gin.Context, sid string, newType StudentRecruitmentCycleType) error {
	tx := db.WithContext(ctx).Model(&StudentRecruitmentCycle{}).Where("student_id = ?", sid).Update("type", newType)
	return tx.Error
}
