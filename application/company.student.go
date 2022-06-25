package application

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spo-iitk/ras-backend/rc"
	"github.com/spo-iitk/ras-backend/util"
)

func getStudentsByEventForCompanyHandler(ctx *gin.Context) {
	cid, err := extractCompanyRCID(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

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

	var proforma Proforma
	err = fetchProforma(ctx, event.ProformaID, &proforma)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if proforma.CompanyRecruitmentCycleID != cid {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Not authorized"})
		return
	}

	students := []EventStudent{}
	err = fetchStudentsByEvent(ctx, eid, &students)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var studentRCIDs []uint
	for _, student := range students {
		studentRCIDs = append(studentRCIDs, student.StudentRecruitmentCycleID)
	}

	var studentRCs []rc.StudentRecruitmentCycle
	err = rc.FetchStudents(ctx, studentRCIDs, &studentRCs)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, studentRCs)
}
