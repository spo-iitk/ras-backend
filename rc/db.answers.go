package rc

import "github.com/gin-gonic/gin"

func fetchStudentAnswers(ctx *gin.Context, sid string, answers *[]RecruitmentCycleQuestionsAnswer) error {
	tx := db.WithContext(ctx).Where("student_recruitment_cycle_id = ?", sid).Find(answers)
	return tx.Error
}

func createStudentAnswer(ctx *gin.Context, answer *RecruitmentCycleQuestionsAnswer) error {
	tx := db.WithContext(ctx).Create(answer)
	return tx.Error
}

func updateStudentAnswer(ctx *gin.Context, answer *RecruitmentCycleQuestionsAnswer) error {
	tx := db.WithContext(ctx).Where("id = ?", answer.ID).Updates(answer)
	return tx.Error
}

func deleteStudentAnswer(ctx *gin.Context, qid string, sid string) error {
	tx := db.WithContext(ctx).Where("recruitment_cycle_question_id = ? AND student_recruitment_cycle_id = ?", qid, sid).Delete(&RecruitmentCycleQuestionsAnswer{})
	return tx.Error
}
