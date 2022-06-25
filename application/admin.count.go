package application

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spo-iitk/ras-backend/util"
)

func getApplicationCountHandler(ctx *gin.Context) {
	var roleCount int
	var recruitmentCount int

	rid, err := util.ParseUint(ctx.Param("rid"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	roleCount, err = fetchRolesCount(ctx, rid)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	recruitmentCount, err = fetchRecruitedCount(ctx, rid)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"roles": roleCount, "recruited": recruitmentCount})
}
