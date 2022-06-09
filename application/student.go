package application

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spo-iitk/ras-backend/util"
)

func getStudentByEventHandler(ctx *gin.Context) {
	eid_string := ctx.Param("eid")
	eid, err := util.ParseUint(eid_string)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	students := []EventStudent{}
	err = fetchStudentsByEvent(ctx, eid, &students)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, students)
}
