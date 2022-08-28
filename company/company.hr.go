package company

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func postNewHRHandler(ctx *gin.Context) {
	var addHRRequest CompanyHR

	err := ctx.ShouldBindJSON(&addHRRequest)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if addHRRequest.CompanyID != 0 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Company ID is not allowed"})
		return
	}

	addHRRequest.CompanyID, err = extractCompanyID(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = addHR(ctx, &addHRRequest)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "Successfully added"})

}

func getCompanyHRHandler(ctx *gin.Context) {
	var HRs []CompanyHR

	cid, err := extractCompanyID(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = getAllHR(ctx, &HRs, cid)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, HRs)
}
