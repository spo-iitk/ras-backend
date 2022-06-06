package rc

import "github.com/gin-gonic/gin"

func putPPOPIO(ctx *gin.Context) {
	sid := ctx.Param("sid")

	err := updateStudentType(ctx, sid, PIOPPO)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"status": "updated student pioppo"})
}
