package application

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spo-iitk/ras-backend/util"
)

type statsResponse struct {
	StudentRecruitmentCycleID uint   `json:"student_recruitment_cycle_id"`
	CompanyName               string `json:"company_name"`
	Role                      string `json:"role"`
	Type                      string `json:"type"`
}

func getStatsHandler(ctx *gin.Context) {
	rid, err := util.ParseUint(ctx.Param("rid"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var stats []statsResponse
	err = getRecruitmentStats(ctx, rid, &stats)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, stats)
}
