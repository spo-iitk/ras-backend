package application

import "github.com/gin-gonic/gin"

func fetchProformaQuestion(ctx *gin.Context, pid uint, questions *[]ApplicationQuestion) error {
	tx := db.WithContext(ctx).Where("proforma_id = ?", pid).Find(&questions)
	return tx.Error
}
func fetchAllAnswers(ctx *gin.Context, pid uint, questionID []uint, answers *[]ApplicationQuestionAnswer) error {

	tx := db.WithContext(ctx).Where("application_question_id IN ?", questionID).Find(answers)
	return tx.Error
}

func updateProformaQuestion(ctx *gin.Context, question *ApplicationQuestion) error {
	tx := db.WithContext(ctx).Where("id = ?", question.ID).Updates(question)
	return tx.Error
}

func createProformaQuestion(ctx *gin.Context, question *ApplicationQuestion) error {
	tx := db.WithContext(ctx).Create(question)
	return tx.Error
}

func deleteProformaQuestion(ctx *gin.Context, qid uint) error {
	tx := db.WithContext(ctx).Where("id = ?", qid).Delete(&ApplicationQuestion{})
	return tx.Error
}

func fetchApplicationQuestionsAnswers(ctx *gin.Context, pid, sid uint, questions *[]getApplicationResponse) error {
	tx := db.WithContext(ctx).Model(&ApplicationQuestion{}).
		Joins("LEFT JOIN application_question_answers ON application_question_answers.application_question_id = application_questions.id AND application_question_answers.student_recruitment_cycle_id = ?", sid).
		Select("application_questions.*, application_question_answers.answer").
		Where("application_questions.proforma_id = ?", pid).
		Find(questions)
	return tx.Error
}
