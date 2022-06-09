package application

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getEventsByStudentHandler(ctx *gin.Context) {
	sid, err := extractStudentRCID(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var events []ProformaEvent
	err = fetchEventsByStudent(ctx, sid, &events)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, events)
}
