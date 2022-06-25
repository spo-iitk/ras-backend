package application

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spo-iitk/ras-backend/util"
)

func getEventsByStudentHandler(ctx *gin.Context) {
	sid, err := extractStudentRCID(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var events []ProformaEvent
	err = fetchEventsByStudent(ctx, sid, &events)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, events)
}

func getEventsByProformaForStudentHandler(ctx *gin.Context) {
	_, err := extractStudentRCID(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Fail to extract student rcid"})
		return
	}

	pid, err := util.ParseUint(ctx.Param("pid"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var events []ProformaEvent
	err = fetchEventsByProforma(ctx, pid, &events)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, events)
}
