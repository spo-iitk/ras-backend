package rc

import "github.com/gin-gonic/gin"

func getCompanyRecruitmentCycle(ctx *gin.Context) {

	var rcs []RecruitmentCycle
	companyID, err := extractCompanyRCID(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}
	err = fetchRCsByCompanyID(ctx, companyID, &rcs)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, rcs)
}
