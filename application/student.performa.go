package application

import (
	"github.com/gin-gonic/gin"
	"github.com/spo-iitk/ras-backend/util"
)

func getProformaByRIDHandler(ctx *gin.Context) {
	rid, err := util.ParseUint(ctx.Param("rid"))
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	var jps []Proforma

	err = fetchProformaByRC(ctx, rid, &jps)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, jps)
}
