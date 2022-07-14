package application

import "github.com/gin-gonic/gin"

func fetchEventsByRC(ctx *gin.Context, rid uint, events *[]getAllEventsByRCResponse) error {
	tx := db.WithContext(ctx).Model(&ProformaEvent{}).
		Joins("JOIN proformas ON proformas.id = proforma_events.proforma_id").
		Where("proformas.deleted_at IS NULL AND proformas.recruitment_cycle_id = ?", rid).
		Order("start_time DESC, proforma_id, sequence").
		Select("proforma_events.*, proformas.company_name, proformas.role").
		Find(events)
	return tx.Error
}

func fetchEvent(ctx *gin.Context, id uint, event *ProformaEvent) error {
	tx := db.WithContext(ctx).Where("id = ?", id).Order("sequence").First(event)
	return tx.Error
}

func fetchEventsByProforma(ctx *gin.Context, pid uint, events *[]ProformaEvent) error {
	tx := db.WithContext(ctx).Where("proforma_id = ?", pid).Order("sequence").Find(events)
	return tx.Error
}

func createEvent(ctx *gin.Context, event *ProformaEvent) error {
	tx := db.WithContext(ctx).Where("proforma_id = ? AND name = ?", event.ProformaID, event.Name).FirstOrCreate(event)
	return tx.Error
}

func updateEvent(ctx *gin.Context, event *ProformaEvent) error {
	tx := db.WithContext(ctx).Where("id = ?", event.ID).Updates(event)
	return tx.Error
}

func deleteEvent(ctx *gin.Context, id uint) error {
	tx := db.WithContext(ctx).Where("id = ?", id).Delete(&ProformaEvent{})
	return tx.Error
}

func fetchEventsByStudent(ctx *gin.Context, sid uint, events *[]ProformaEvent) error {
	tx := db.WithContext(ctx).
		Joins("NATURAL JOIN event_students").
		Where("event_students.student_recruitment_cycle_id = ?", sid).
		Order("start_time DESC, proforma_id, sequence").Find(events)
	return tx.Error
}
