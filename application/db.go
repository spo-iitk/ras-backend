package application

import (
	"github.com/gin-gonic/gin"
)

func fetchEventsByRC(ctx *gin.Context, rid uint, events *[]JobProformaEvent) error {
	tx := db.WithContext(ctx).Joins("job_proforma", db.Where(&JobProforma{RecruitmentCycleID: rid})).Find(events)
	return tx.Error
}

func firstOrCreateEmptyPerfoma(ctx *gin.Context, jp *JobProforma) error {
	tx := db.WithContext(ctx).Where("company_recruitment_cycle_id = ?", jp.CompanyRecruitmentCycleID).FirstOrCreate(jp)
	return tx.Error
}

func createEvent(ctx *gin.Context, event *JobProformaEvent) error {
	tx := db.WithContext(ctx).Create(event)
	return tx.Error
}

// func createStudentEvent(ctx *gin.Context, eventStudent *EventStudent) error {
// 	tx := db.WithContext(ctx).Create(eventStudent)
// 	return tx.Error
// }

func createStudentEvents(ctx *gin.Context, eventStudents *[]EventStudent) error {
	tx := db.WithContext(ctx).Create(eventStudents)
	return tx.Error
}
