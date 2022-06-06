package rc

import "github.com/gin-gonic/gin"

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

func getRC(ctx *gin.Context) {
	id := ctx.Param("rid")
	var rc RecruitmentCycle
	err := fetchRC(ctx, id, &rc)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, rc)
}
