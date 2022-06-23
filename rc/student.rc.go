package rc

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spo-iitk/ras-backend/middleware"
	"github.com/spo-iitk/ras-backend/util"
)

func getStudentRC(ctx *gin.Context) {
	email := middleware.GetUserID(ctx)

	var rcs []RecruitmentCycle
	err := fetchRCsByStudent(ctx, email, &rcs)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, rcs)
}

func getStudent(ctx *gin.Context) {
	rid, err := util.ParseUint(ctx.Param("rid"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	email := middleware.GetUserID(ctx)
	var student StudentRecruitmentCycle

	err = fetchStudent(ctx, email, rid, &student)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, student)
}
