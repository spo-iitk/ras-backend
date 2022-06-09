package application

import (
	"github.com/gin-gonic/gin"
	"github.com/spo-iitk/ras-backend/util"
)

func getStats(ctx *gin.Context) {
	rid_string := ctx.Param("rid")
	rid, err := util.ParseUint(rid_string)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var stats []EventStudent
	err = getRecruitmentStats(ctx, rid, &stats)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, stats)
}
