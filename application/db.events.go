package application

import "github.com/gin-gonic/gin"

func fetchEventsByRC(ctx *gin.Context, rid uint, events *[]ProformaEvent) error {
	tx := db.WithContext(ctx).Joins("proforma", db.Where(&Proforma{RecruitmentCycleID: rid})).Find(events)
	return tx.Error
}

func fetchEventByID(ctx *gin.Context, id uint, event *ProformaEvent) error {
	tx := db.WithContext(ctx).Where("id = ?", id).First(event)
	return tx.Error
}

func fetchEventsByPID(ctx *gin.Context, pid uint, events *[]ProformaEvent) error {
	tx := db.WithContext(ctx).Where("proforma_id = ?", pid).Find(events)
	return tx.Error
}

func createEvent(ctx *gin.Context, event *ProformaEvent) error {
	tx := db.WithContext(ctx).Create(event)
	return tx.Error
}

func updateEvent(ctx *gin.Context, event *ProformaEvent) error {
	tx := db.WithContext(ctx).Where("id = ?", event.ID).Updates(event)
	return tx.Error
}

func deleteEvent(ctx *gin.Context, id uint) error {
	tx := db.WithContext(ctx).Where("id = ?", id).Delete(ProformaEvent{})
	return tx.Error
}

func fetchEventsByStudent(ctx *gin.Context, sid uint, events *[]ProformaEvent) error {
	tx := db.WithContext(ctx).Joins("event_student", db.Where(&EventStudent{StudentRecruitmentCycleID: sid})).Find(events)
	return tx.Error
}
