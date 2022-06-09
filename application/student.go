package application

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spo-iitk/ras-backend/rc"
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

type postStudentsByEventRequest struct {
	EventID uint     `json:"eventID" binding:"required"`
	Emails  []string `json:"emails" binding:"required"`
}

func postStudentsByEventHandler(ctx *gin.Context) {
	rid_string := ctx.Param("rid")
	rid, err := util.ParseUint(rid_string)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var req postStudentsByEventRequest
	err = ctx.BindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	srcIDs, err := rc.FetchStudentRCIDs(ctx, rid, req.Emails)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	students := []EventStudent{}
	for _, srcID := range srcIDs {
		students = append(students, EventStudent{
			JobProformaEventID:        req.EventID,
			StudentRecruitmentCycleID: srcID,
		})
	}

	err = createEventStudents(ctx, &students)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "Added successfully"})
}
