package application

import "github.com/gin-gonic/gin"

func createEventStudent(ctx *gin.Context, eventStudent *EventStudent) error {
	tx := db.WithContext(ctx).FirstOrCreate(eventStudent)
	return tx.Error
}

func createEventStudents(ctx *gin.Context, eventStudents *[]EventStudent) error {
	tx := db.WithContext(ctx).Create(eventStudents)
	return tx.Error
}

func getRecruitmentStats(ctx *gin.Context, rid uint, stats *[]EventStudent) error {
	tx := db.WithContext(ctx).Joins("proforma_event", db.Where("name IN", []EventType{Recruited, PIOPPOACCEPTED})).Where("recruitment_cycle_id = ?", rid).Find(stats)
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
	tx := db.WithContext(ctx).Model(&EventStudent{}).Where("student_recruitment_cycle_id = ?", sid).Group("company_recruitment_cycle_id").Count(&count)
	return int(count), tx.Error
}
