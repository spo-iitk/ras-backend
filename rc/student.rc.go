package rc

import (
	"github.com/gin-gonic/gin"
	"github.com/spo-iitk/ras-backend/middleware"
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
	rid := ctx.Param("rid")
	email := middleware.GetUserID(ctx)
	var student StudentRecruitmentCycle

	err := fetchStudent(ctx, email, rid, &student)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, student)
}
