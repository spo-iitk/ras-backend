package application

import "github.com/gin-gonic/gin"

func getEventsByRC(ctx *gin.Context, rid uint, events *[]JobPerformaEvent) error {
	tx := db.WithContext(ctx).Joins("job_performa", db.Where(&JobProforma{RecruitmentCycleID: rid})).Find(events)
	return tx.Error
}
