package application

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getApplicationCount(ctx *gin.Context) {
	var roleCount int
	var PPOPPIOCount int

	rid, err := strconv.ParseUint(ctx.Param("rid"), 10, 64)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	roleCount, err = getRolesCount(ctx, uint(rid))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	PPOPPIOCount, err = getPPOPIOCount(ctx, uint(rid))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"roles": roleCount, "PPO-PIO": PPOPPIOCount})

}
