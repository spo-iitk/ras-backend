package application

import "github.com/gin-gonic/gin"

func fetchProformaQuestion(ctx *gin.Context, pid uint, questions *[]JobApplicationQuestion) error {
	tx := db.WithContext(ctx).Where("proforma_id = ?", pid).Find(&questions)
	return tx.Error
}

func fetchProformaQuestionByID(ctx *gin.Context, qid uint, question *JobApplicationQuestion) error {
	tx := db.WithContext(ctx).Where("id = ?", qid).First(question)
	return tx.Error
}

func updateProformaQuestion(ctx *gin.Context, question *JobApplicationQuestion) error {
	tx := db.WithContext(ctx).Where("id = ?", question.ID).Updates(question)
	return tx.Error
}

func createProformaQuestion(ctx *gin.Context, question *JobApplicationQuestion) error {
	tx := db.WithContext(ctx).FirstOrCreate(question)
	return tx.Error
}
