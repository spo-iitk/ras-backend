package application

import "github.com/gin-gonic/gin"

func createEventStudents(ctx *gin.Context, eventStudents *[]EventStudent) error {
	tx := db.WithContext(ctx).Create(eventStudents)
	return tx.Error
}

func getRecruitmentStats(ctx *gin.Context, rid uint, stats *[]statsResponse) error {
	tx := db.WithContext(ctx).Model(&EventStudent{}).
		Joins("JOIN proforma_events ON proforma_events.name IN ? AND proforma_events.id = event_students.proforma_event_id", []EventType{Recruited, PIOPPOACCEPTED}).
		Joins("JOIN proformas ON proformas.id = proforma_events.proforma_id AND proformas.recruitment_cycle_id = ?", rid).
		Select("event_students.student_recruitment_cycle_id, proformas.company_name, proformas.profile ,proforma_events.name as type").
		Order("event_students.student_recruitment_cycle_id").
		Find(stats)
	return tx.Error
}

func fetchStudentRCIDByEvents(ctx *gin.Context, eventID uint) ([]uint, error) {
	var ids []uint
	tx := db.WithContext(ctx).Model(&EventStudent{}).Where("proforma_event_id = ?", eventID).Pluck("student_recruitment_cycle_id", &ids)
	return ids, tx.Error
}

func fetchStudentsByEvent(ctx *gin.Context, eventID uint, students *[]EventStudent) error {
	tx := db.WithContext(ctx).Where("proforma_event_id = ?", eventID).Find(students)
	return tx.Error
}

func getCurrentApplicationCount(ctx *gin.Context, sid uint) (int, error) {
	var count int64
	tx := db.WithContext(ctx).Model(&EventStudent{}).
		Where("student_recruitment_cycle_id = ?", sid).
		Group("company_recruitment_cycle_id").Count(&count)
	return int(count), tx.Error
}

func deleteStudentFromEvent(ctx *gin.Context, eventID, studentID uint) error {
	tx := db.WithContext(ctx).Where("proforma_event_id = ? AND student_recruitment_cycle_id = ?", eventID, studentID).Delete(&EventStudent{})
	return tx.Error
}

func deleteAllStudentsFromEvent(ctx *gin.Context, eventID uint) error {
	tx := db.WithContext(ctx).Where("proforma_event_id = ?", eventID).Delete(&EventStudent{})
	return tx.Error
}

func getStudentIDByEventID(ctx *gin.Context, eid uint) ([]uint, error) {
	var studentsIDs []uint
	tx := db.WithContext(ctx).Where("proforma_event_id = ?", eid).Model(&EventStudent{}).
		Pluck("student_recruitment_cycle_id", &studentsIDs)
	return studentsIDs, tx.Error
}
