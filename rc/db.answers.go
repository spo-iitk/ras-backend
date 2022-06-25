package rc

import "github.com/gin-gonic/gin"

func createStudentAnswer(ctx *gin.Context, answer *RecruitmentCycleQuestionsAnswer) error {
	tx := db.WithContext(ctx).
		Where(
			"recruitment_cycle_question_id = ? AND student_recruitment_cycle_id = ?",
			answer.RecruitmentCycleQuestionID,
			answer.StudentRecruitmentCycleID,
		).FirstOrCreate(answer)
	return tx.Error
}
