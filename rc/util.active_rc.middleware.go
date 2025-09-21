package rc

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func checkAdminAccessToRC() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		roleID := ctx.GetInt("roleID")
		if roleID > 101 && !checkIsActiveRC(ctx) {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}
		ctx.Next()
	}
}

func checkIsActiveRC(ctx *gin.Context) bool {
	id := ctx.Param("rid")
	var rc RecruitmentCycle
	err := fetchRC(ctx, id, &rc)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return false
	}
	return rc.IsActive
}
