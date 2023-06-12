package company

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func getAllCompanyHRsHandler(ctx *gin.Context) {
	var companyHRs []CompanyHR

	err := getAllHRUserDB(ctx, &companyHRs)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, companyHRs)
}
