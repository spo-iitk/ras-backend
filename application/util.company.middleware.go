package application

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spo-iitk/ras-backend/rc"
	"github.com/spo-iitk/ras-backend/util"
)

func ensureCompany() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		rid, err := util.ParseUint(ctx.Param("rid"))
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if !rc.IsRCActive(ctx, rid) {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "RC not active"})
			return
		}

		companyID, err := extractCompanyID(ctx)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx.Set("companyID", companyID)

		crcid, err := rc.FetchCompanyRCID(ctx, rid, companyID)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx.Set("companyRCID", crcid)

		ctx.Next()
	}
}

// func getCompanyID(ctx *gin.Context) uint {
// 	return uint(ctx.GetInt("companyID"))
// }

func getCompanyRCID(ctx *gin.Context) uint {
	return uint(ctx.GetInt("companyRCID"))
}
