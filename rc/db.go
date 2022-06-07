package rc

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func fetchAllRCs(ctx *gin.Context, rc *[]RecruitmentCycle) error {
	tx := db.WithContext(ctx).Find(rc)
	return tx.Error
}

func fetchRCsByStudent(ctx *gin.Context, email string, rcs *[]RecruitmentCycle) error {
	tx := db.WithContext(ctx).Joins("student_recruitment_cycle", db.Where(&StudentRecruitmentCycle{Email: email})).Find(&rcs)
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

func fetchAllNotices(ctx *gin.Context, rid string, notices *[]Notice) error {
	tx := db.WithContext(ctx).Where("recruitment_cycle_id = ?", rid).Find(notices)
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
	tx := db.WithContext(ctx).Where("recruitment_cycle_id = ?", rid).Find(companies)
	return tx.Error
}

func fetchCompany(ctx *gin.Context, rid string, cid string, company *CompanyRecruitmentCycle) error {
	tx := db.WithContext(ctx).Where("recruitment_cycle_id = ? AND company_id = ?", rid, cid).First(company)
	return tx.Error
}

func createCompany(ctx *gin.Context, company *CompanyRecruitmentCycle) error {
	tx := db.WithContext(ctx).Create(company)
	return tx.Error
}

func fetchAllStudents(ctx *gin.Context, rid string, students *[]StudentRecruitmentCycle) error {
	tx := db.WithContext(ctx).Where("recruitment_cycle_id = ?", rid).Find(students)
	return tx.Error
}

func fetchStudent(ctx *gin.Context, email string, rid string, student *StudentRecruitmentCycle) error {
	tx := db.WithContext(ctx).Where("email = ? AND recruitment_cycle_id = ?", email, rid).First(student)
	return tx.Error
}

func updateStudent(ctx *gin.Context, student *StudentRecruitmentCycle) error {
	tx := db.WithContext(ctx).Save(student)
	return tx.Error
}

func createStudents(ctx *gin.Context, students *[]StudentRecruitmentCycle) error {
	tx := db.WithContext(ctx).Create(students)
	return tx.Error
}

func fetchStudentQuestions(ctx *gin.Context, rid string, questions *[]RecruitmentCycleQuestion) error {
	tx := db.WithContext(ctx).Where("recruitment_cycle_id = ?", rid).Find(questions)
	return tx.Error
}

func createStudentQuestion(ctx *gin.Context, question *RecruitmentCycleQuestion) error {
	tx := db.WithContext(ctx).Create(question)
	return tx.Error
}

func updateStudentQuestion(ctx *gin.Context, question *RecruitmentCycleQuestion) error {
	tx := db.WithContext(ctx).Save(question)
	return tx.Error
}

func deleteStudentQuestion(ctx *gin.Context, qid string) error {
	tx := db.WithContext(ctx).Where("id = ?", qid).Delete(RecruitmentCycleQuestion{})
	return tx.Error
}

func fetchStudentAnswers(ctx *gin.Context, sid string, answers *[]RecruitmentCycleQuestionsAnswer) error {
	tx := db.WithContext(ctx).Where("student_recruitment_cycle_id = ?", sid).Find(answers)
	return tx.Error
}

func createStudentAnswer(ctx *gin.Context, answer *RecruitmentCycleQuestionsAnswer) error {
	tx := db.WithContext(ctx).Create(answer)
	return tx.Error
}

func updateStudentAnswer(ctx *gin.Context, answer *RecruitmentCycleQuestionsAnswer) error {
	tx := db.WithContext(ctx).Save(answer)
	return tx.Error
}

func deleteStudentAnswer(ctx *gin.Context, qid string, sid string) error {
	tx := db.WithContext(ctx).Where("recruitment_cycle_question_id = ? AND student_recruitment_cycle_id = ?", qid, sid).Delete(RecruitmentCycleQuestionsAnswer{})
	return tx.Error
}

func updateStudentType(ctx *gin.Context, r *pioppoRequest, newType StudentRecruitmentCycleType) error {
	cid := r.cid
	emails := r.email

	var c CompanyRecruitmentCycle
	tx := db.WithContext(ctx).Where("id = ?", cid).First(c)
	if tx.Error != nil {
		return tx.Error
	}

	tx = db.WithContext(ctx).Model(StudentRecruitmentCycle{}).Where("recruitment_cycle_id = ? AND email IN ?", c.RecruitmentCycleID, emails).Updates(
		StudentRecruitmentCycle{
			Type:     newType,
			IsFrozen: true,
			Comment:  "PIO/PPO by " + c.CompanyName,
		})
	return tx.Error
}
