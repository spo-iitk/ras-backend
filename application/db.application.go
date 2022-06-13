package application

import "github.com/gin-gonic/gin"

func fetchApplicationEventID(ctx *gin.Context, pid uint) (uint, error) {
	var event ProformaEvent
	tx := db.WithContext(ctx).Where("proforma_id = ? AND name = ?", pid, ApplicationSubmitted).First(&event)
	return event.ID, tx.Error
}

func deleteApplication(ctx *gin.Context, pid uint, sid uint) error {
	tx := db.WithContext(ctx).Where("proforma_event_id = ? AND student_recruitment_cycle_id = ?", pid, sid).Delete(&EventStudent{})
	return tx.Error
}

func fetchApplicationCount(ctx *gin.Context, sid uint) (bool, error) {
	var count int64
	tx := db.WithContext(ctx).Model(&EventStudent{}).Distinct().Where("student_recruitment_cycle_id = ?", sid).Count(&count)
	return false, tx.Error
}
