package application

import (
	"github.com/gin-gonic/gin"
)

func fetchApplicationEventID(ctx *gin.Context, pid uint) (uint, error) {
	var event ProformaEvent
	tx := db.WithContext(ctx).Where("proforma_id = ? AND name = ?", pid, ApplicationSubmitted).First(&event)
	return event.ID, tx.Error
}

func deleteApplication(ctx *gin.Context, pid uint, sid uint) error {
	var questions []getApplicationResponse
	err := fetchApplicationQuestionsAnswers(ctx, pid, sid, &questions)
	if err != nil {
		return err
	}

	var qid []uint
	for _, question := range questions {
		qid = append(qid, question.ID)
	}

	tx := db.WithContext(ctx).Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	if err := tx.Where("proforma_event_id = ? AND student_recruitment_cycle_id = ?", pid, sid).Delete(&EventStudent{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Where("application_question_id IN ?", qid).Delete(&ApplicationQuestionAnswer{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Where("proforma_id = ?", pid).Delete(&ApplicationResume{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func createApplication(ctx *gin.Context, application *EventStudent, answers *[]ApplicationQuestionAnswer, resume *ApplicationResume) error {
	tx := db.WithContext(ctx).Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	err := tx.
		Where(
			"proforma_event_id = ? AND student_recruitment_cycle_id = ?",
			application.ProformaEventID,
			application.StudentRecruitmentCycleID).
		FirstOrCreate(application).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Create(answers).Error; err != nil {
		tx.Rollback()
		return err
	}

	err = tx.
		Where("proforma_id = ? AND student_recruitment_cycle_id = ?", resume.ProformaID, resume.ResumeID).
		FirstOrCreate(resume).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func fetchApplicantDetails(ctx *gin.Context, pid uint, students *[]ApplicantsByRole) error {
	tx := db.WithContext(ctx).Model(&EventStudent{}).
		Joins("JOIN proforma_events ON proforma_events.id = event_students.proforma_event_id").
		Joins("JOIN application_resumes ON application_resumes.proforma_id = proforma_events.id").
		Where("proforma_events.proforma_id = ? ", pid).
		Select("event_students.student_recruitment_cycle_id as student_id, application_resumes.resume as resume_link, proforma_events.sequence as status").
		// Group("event_students.student_recruitment_cycle_id").
		Scan(students)

	return tx.Error
}
