package application

import (
	"database/sql"
	"log"

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

	var events []ProformaEvent
	err = fetchEventsByProforma(ctx, pid, &events)
	if err != nil {
		return err
	}

	log.Println("Failed to delete application")
	var eid []uint
	for _, event := range events {
		eid = append(eid, event.ID)
	}

	log.Println("Failed to delete applicationsadfg")
	tx := db.WithContext(ctx).Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	log.Println("Failed sfdgto delete applicationsadfg")
	if err := tx.Error; err != nil {
		return err
	}
	log.Println("Failed sfdgto delete applicationsadfg")

	if err := tx.Where("proforma_event_id IN ? AND student_recruitment_cycle_id = ?", eid, sid).Delete(&EventStudent{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	log.Println("Failed sfgvjhbknlm;,dgto delete applicationsadfg")

	if len(qid) > 0 {
		if err := tx.Where("application_question_id IN ? AND student_recruitment_cycle_id = ?", qid, sid).Delete(&ApplicationQuestionAnswer{}).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	log.Println("Failelnkm;bhvbjnd sfgvjhbknlm;,dgto delete applicationsadfg")
	if err := tx.Where("proforma_id = ? AND student_recruitment_cycle_id = ?", pid, sid).Delete(&ApplicationResume{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
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

	if len(*answers) > 0 {
		err = tx.Create(answers).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	err = tx.
		Where("proforma_id = ? AND student_recruitment_cycle_id = ?", resume.ProformaID, resume.StudentRecruitmentCycleID).
		FirstOrCreate(resume).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func fetchApplicantDetails(ctx *gin.Context, pid uint, students *[]ApplicantsByRole) error {
	query := `
	SELECT
		*
	FROM
		(
			SELECT
				event_students.student_recruitment_cycle_id AS student_rc_id,
				application_resumes.resume AS resume_link,
				proforma_events.sequence AS status,
				proforma_events.name AS name,
				application_resumes.proforma_id
			FROM
				event_students
				JOIN proforma_events ON proforma_events.id = event_students.proforma_event_id
				JOIN application_resumes ON application_resumes.proforma_id = proforma_events.proforma_id
				AND event_students.student_recruitment_cycle_id = application_resumes.student_recruitment_cycle_id
			WHERE
				proforma_events.proforma_id = @pid
				AND event_students.deleted_at IS NULL
				AND application_resumes.deleted_at IS NULL
				AND proforma_events.deleted_at IS NULL
		) mulstatus NATURAL
		JOIN (
			SELECT
				student_rc_id,
				Max(status) AS status,
				proforma_id
			FROM
				(
					SELECT
						event_students.student_recruitment_cycle_id AS student_rc_id,
						application_resumes.resume AS resume_link,
						proforma_events.sequence AS status,
						proforma_events.name AS name,
						application_resumes.proforma_id
					FROM
						event_students
						JOIN proforma_events ON proforma_events.id = event_students.proforma_event_id
						JOIN application_resumes ON application_resumes.proforma_id = proforma_events.proforma_id
						AND event_students.student_recruitment_cycle_id = application_resumes.student_recruitment_cycle_id
					WHERE
						proforma_events.proforma_id = @pid
						AND event_students.deleted_at IS NULL
						AND application_resumes.deleted_at IS NULL
						AND proforma_events.deleted_at IS NULL
				) mul
			GROUP BY
				student_rc_id,
				proforma_id
		) ms
		ORDER BY
		student_rc_id`
	tx := db.WithContext(ctx).
		Raw(query, sql.Named("pid", pid)).
		Scan(students)

	return tx.Error
}

func fetchApplications(ctx *gin.Context, sid uint, response *[]ViewApplicationsBySIDResponse) error {
	tx := db.WithContext(ctx).Model(&EventStudent{}).
		Joins("JOIN proforma_events ON proforma_events.id = event_students.proforma_event_id AND proforma_events.deleted_at IS NULL").
		Joins("JOIN application_resumes ON application_resumes.student_recruitment_cycle_id = event_students.student_recruitment_cycle_id AND application_resumes.deleted_at IS NULL AND application_resumes.proforma_id = proforma_events.proforma_id").
		Joins("JOIN proformas ON proformas.id = proforma_events.proforma_id AND proformas.deleted_at IS NULL").
		Where("event_students.student_recruitment_cycle_id = ? AND event_students.deleted_at IS NULL AND proforma_events.name = ?", sid, ApplicationSubmitted).
		Distinct("proformas.ID, proformas.company_name, proformas.role, proformas.deadline, application_resumes.resume_id, application_resumes.resume").
		Scan(response)

	return tx.Error
}

func fetchAdminApplications(ctx *gin.Context, sid uint, response *[]ViewApplicationsBySIDAdminResponse) error {
	tx := db.WithContext(ctx).Model(&EventStudent{}).
		Joins("JOIN proforma_events ON proforma_events.id = event_students.proforma_event_id AND proforma_events.deleted_at IS NULL").
		Joins("JOIN application_resumes ON application_resumes.student_recruitment_cycle_id = event_students.student_recruitment_cycle_id AND application_resumes.deleted_at IS NULL AND application_resumes.proforma_id = proforma_events.proforma_id").
		Joins("JOIN proformas ON proformas.id = proforma_events.proforma_id AND proformas.deleted_at IS NULL").
		Where("event_students.student_recruitment_cycle_id = ? AND event_students.deleted_at IS NULL AND proforma_events.name = ?", sid, ApplicationSubmitted).
		Distinct("proformas.ID, proformas.company_name, proformas.role, proformas.deadline, application_resumes.resume_id, application_resumes.resume, event_students.created_at as applied_on").
		Scan(response)

	return tx.Error
}
