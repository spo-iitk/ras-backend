package rc

import (
	"github.com/gin-gonic/gin"
	"github.com/spo-iitk/ras-backend/middleware"
)

func getAllRC(ctx *gin.Context) {
	var rc []RecruitmentCycle
	err := fetchAllRCs(ctx, &rc)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, rc)
}

func postRC(ctx *gin.Context) {
	var rc RecruitmentCycle
	err := ctx.BindJSON(&rc)
	if err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}
	err = createRC(ctx, &rc)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}
	id := gin.H{"id": rc.ID}
	ctx.JSON(201, gin.H{"status": "created", "data": id})
}

//! TODO: Add more response data
type getRCResponse struct {
	RecruitmentCycle
}

func getRC(ctx *gin.Context) {
	id := ctx.Param("rid")
	var rc RecruitmentCycle
	err := fetchRC(ctx, id, &rc)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, getRCResponse{rc})
}

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

func getCompanyRecruitmentCycle(ctx *gin.Context) {
	// email := middleware.GetUserID(ctx)

	var rcs []RecruitmentCycle
	companyID := uint(5) //! TODO get from company
	err := fetchRCsByCompanyID(ctx, companyID, &rcs)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, rcs)
}
