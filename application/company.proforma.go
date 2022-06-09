package application

import (
	"github.com/gin-gonic/gin"
)

func getProformaForCompanyHandler(ctx *gin.Context) {
	cid, err := extractCompanyRCID(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}

	var jps []Proforma

	err = fetchProformaByCompanyRC(ctx, cid, &jps)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, jps)
}
