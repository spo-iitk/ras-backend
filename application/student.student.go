package application

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spo-iitk/ras-backend/rc"
	"github.com/spo-iitk/ras-backend/util"
)

func getStudentsByEventForStudentHandler(ctx *gin.Context) {
	eid, err := util.ParseUint(ctx.Param("eid"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var event ProformaEvent
	err = fetchEvent(ctx, eid, &event)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if event.StartTime == 0 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Event has not started yet"})
		return
	}

	var students []EventStudent
	err = fetchStudentsByEvent(ctx, eid, &students)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var studentRCIDs []uint
	for _, student := range students {
		studentRCIDs = append(studentRCIDs, student.StudentRecruitmentCycleID)
	}

	var studentRCs []rc.StudentRecruitmentCycle
	err = rc.FetchStudents(ctx, studentRCIDs, &studentRCs)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i := range studentRCs {
		studentRCs[i].StudentID = 0
		studentRCs[i].RecruitmentCycleID = 0
		studentRCs[i].CPI = 0
		studentRCs[i].Type = ""
		studentRCs[i].IsFrozen = false
		studentRCs[i].IsVerified = false
		studentRCs[i].Comment = ""
	}

	ctx.JSON(http.StatusOK, studentRCs)
}
