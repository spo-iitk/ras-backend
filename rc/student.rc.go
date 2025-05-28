package rc

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spo-iitk/ras-backend/middleware"
	"github.com/spo-iitk/ras-backend/util"
)

func getStudentRCHandler(ctx *gin.Context) {
	email := middleware.GetUserID(ctx)

	var rcs []RecruitmentCycle
	err := fetchRCsByStudent(ctx, email, &rcs)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, rcs)
}

func studentWhoamiHandler(ctx *gin.Context) {
	rid, err := util.ParseUint(ctx.Param("rid"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	email := middleware.GetUserID(ctx)
	var student StudentRecruitmentCycle

	err = fetchStudentByEmailAndRC(ctx, email, rid, &student)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, student)
}
