package application

import (
	"database/sql"

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
	tx := db.WithContext(ctx).
		Raw("SELECT * FROM (SELECT event_students.student_recruitment_cycle_id AS student_rc_id, application_resumes.resume AS resume_link, proforma_events.sequence AS status, proforma_events.name AS name, application_resumes.proforma_id FROM event_students JOIN proforma_events ON proforma_events.id = event_students.proforma_event_id JOIN application_resumes ON application_resumes.proforma_id = proforma_events.proforma_id WHERE proforma_events.proforma_id = @pid AND event_students.deleted_at IS NULL) mulstatus NATURAL JOIN (SELECT student_rc_id, Max(status) AS status, proforma_id FROM (SELECT event_students.student_recruitment_cycle_id AS student_rc_id, application_resumes.resume AS resume_link, proforma_events.sequence AS status, proforma_events.name AS name, application_resumes.proforma_id FROM event_students JOIN proforma_events ON proforma_events.id = event_students.proforma_event_id JOIN application_resumes ON application_resumes.proforma_id = proforma_events.proforma_id WHERE proforma_events.proforma_id = @pid AND event_students.deleted_at IS NULL) mul GROUP BY student_rc_id, proforma_id) ms", sql.Named("pid", pid)).
		Scan(students)

	return tx.Error
}
