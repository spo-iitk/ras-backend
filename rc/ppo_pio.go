package rc

import "github.com/gin-gonic/gin"

type pioppoRequest struct {
	cid   string
	email []string
}

func postPPOPIO(ctx *gin.Context) {
	var req pioppoRequest

	err := ctx.BindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	err = updateStudentType(ctx, &req, PIOPPO)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"status": "updated student pioppo"})
}
