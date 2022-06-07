package company

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getAllCompaniesHandler(ctx *gin.Context) {
	var companies []Company

	err := getAllCompanies(ctx, &companies)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, gin.H{"data": companies})
}

func getCompanyHandler(ctx *gin.Context) {
	var company Company

	cid, err := strconv.ParseUint(ctx.Param("cid"), 10, 64)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = getCompany(ctx, &company, uint(cid))

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": company})
}
