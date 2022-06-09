package application

import "github.com/gin-gonic/gin"

func fetchCoordinatorsByEvent(ctx *gin.Context, eventID uint, students *[]EventCoordinator) error {
	tx := db.WithContext(ctx).Where("proforma_event_id = ?", eventID).Find(students)
	return tx.Error
}

func createEventCoordinator(ctx *gin.Context, eventCoordinator *EventCoordinator) error {
	tx := db.WithContext(ctx).FirstOrCreate(eventCoordinator)
	return tx.Error
}



