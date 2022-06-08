package application

import "github.com/gin-gonic/gin"

func fetchPerformaByCompanyRC(ctx *gin.Context, cid uint, jps *[]JobProforma) error {
	tx := db.WithContext(ctx).Where("company_recruitment_cycle_id = ?", cid).Find(jps)
	return tx.Error
}

func fetchPerformaByRC(ctx *gin.Context, rid uint, jps *[]JobProforma) error {
	tx := db.WithContext(ctx).Where("recruitment_cycle_id = ?", rid).Find(jps)
	return tx.Error
}

func fetchJobPerforma(ctx *gin.Context, pid uint, jp *JobProforma) error {
	tx := db.WithContext(ctx).Where("id = ?", pid).First(jp)
	return tx.Error
}

func createJobPerforma(ctx *gin.Context, jp *JobProforma) error {
	tx := db.WithContext(ctx).Create(jp)
	return tx.Error
}

func updateJobPerforma(ctx *gin.Context, jp *JobProforma) error {
	tx := db.WithContext(ctx).Where("id = ?", jp.ID).Updates(jp)
	return tx.Error
}

func deleteJobPerforma(ctx *gin.Context, pid uint) error {
	tx := db.WithContext(ctx).Where("id = ?", pid).Delete(JobProforma{})
	return tx.Error
}

func fetchPerformaQuestion(ctx *gin.Context, pid uint, questions *[]JobApplicationQuestion) error {
	tx := db.WithContext(ctx).Where("job_proforma_id = ?", pid).Find(&questions)
	return tx.Error
}

func fetchPerformaQuestionByID(ctx *gin.Context, qid uint, question *JobApplicationQuestion) error {
	tx := db.WithContext(ctx).Where("id = ?", qid).First(question)
	return tx.Error
}

func updatePerformaQuestion(ctx *gin.Context, question *JobApplicationQuestion) error {
	tx := db.WithContext(ctx).Where("id = ?", question.ID).Updates(question)
	return tx.Error
}

func createPerformaQuestion(ctx *gin.Context, question *JobApplicationQuestion) error {
	tx := db.WithContext(ctx).Create(question)
	return tx.Error
}
