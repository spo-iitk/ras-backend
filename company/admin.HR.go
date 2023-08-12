package company

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func getAllHRHandler(ctx *gin.Context) {
	var HRs []CompanyHR

	cid, err := strconv.ParseUint(ctx.Param("cid"), 10, 32)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = getAllHR(ctx, &HRs, uint(cid))

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, HRs)
}

func deleteHRHandler(ctx *gin.Context) {

	hrid, err := strconv.ParseUint(ctx.Param("hrid"), 10, 32)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = deleteHR(ctx, uint(hrid))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	logrus.Infof("An HR with id %d is deleted", hrid)

	ctx.JSON(http.StatusOK, gin.H{"status": "Successfully deleted"})
}

func addHRHandler(ctx *gin.Context) {
	var addHRRequest CompanyHR

	if err := ctx.ShouldBindJSON(&addHRRequest); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := addHR(ctx, &addHRRequest)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	logrus.Infof("An HR %s is added with id %d", addHRRequest.Name, addHRRequest.ID)

	ctx.JSON(http.StatusOK, gin.H{"status": "Successfully added"})

}

func getInactiveHRsHandler(ctx *gin.Context) {
	var HR []CompanyHR

	cid, err := strconv.ParseUint(ctx.Param("cid"), 10, 32)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = getInactiveHR(ctx, uint(cid), &HR)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, HR)
}
