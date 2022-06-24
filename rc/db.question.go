package rc

import "github.com/gin-gonic/gin"

func fetchStudentQuestions(ctx *gin.Context, rid string, questions *[]RecruitmentCycleQuestion) error {
	tx := db.WithContext(ctx).Where("recruitment_cycle_id = ?", rid).Find(questions)
	return tx.Error
}

func createStudentQuestion(ctx *gin.Context, question *RecruitmentCycleQuestion) error {
	tx := db.WithContext(ctx).Create(question)
	return tx.Error
}

func updateStudentQuestion(ctx *gin.Context, question *RecruitmentCycleQuestion) (bool, error) {
	tx := db.WithContext(ctx).Where("id =?", question.ID).Updates(question)
	return tx.RowsAffected > 0, tx.Error
}

func deleteStudentQuestion(ctx *gin.Context, qid string) error {
	tx := db.WithContext(ctx).Where("id = ?", qid).Delete(&RecruitmentCycleQuestion{})
	return tx.Error
}

func fetchStudentQuestionsAnswers(ctx *gin.Context, rid, sid uint, questions *[]getStudentEnrollmentResponse) error {
	tx := db.WithContext(ctx).Model(&RecruitmentCycleQuestion{}).
		Joins("LEFT JOIN recruitment_cycle_questions_answers ON recruitment_cycle_questions_answers.recruitment_cycle_question_id = recruitment_cycle_questions.id").
		Select("recruitment_cycle_questions.*, recruitment_cycle_questions_answers.answer").
		Where(
			"recruitment_cycle_questions.recruitment_cycle_id = ? AND (recruitment_cycle_questions_answers.student_recruitment_cycle_id = ? OR recruitment_cycle_questions_answers.student_recruitment_cycle_id IS NULL)",
			rid, sid).
		Find(questions)
	return tx.Error
}
