package company

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func getAllHRHandler(ctx *gin.Context) {
	var HRs []CompanyHR

	cid, err := strconv.ParseUint(ctx.Param("cid"), 10, 64)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = getAllHR(ctx, &HRs, uint(cid))

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": HRs})
}

func getHRHandler(ctx *gin.Context) {
	var getHRRequest CompanyHR

	id, err := strconv.ParseUint(ctx.Param("hrid"), 10, 64)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = getHR(ctx, &getHRRequest, uint(id))

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": getHRRequest})
}

func deleteHRHandler(ctx *gin.Context) {

	// var deleteHRRequest CompanyHR

	id, err := strconv.ParseUint(ctx.Param("hrid"), 10, 64)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = deleteHR(ctx, uint(id))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	logrus.Infof("An HR with id %d is deleted", id)

	ctx.JSON(http.StatusOK, gin.H{"status": "Successfully deleted"})
}

func addHRHandler(ctx *gin.Context) {
	var addHRRequest CompanyHR

	cid, err := strconv.ParseUint(ctx.Param("cid"), 10, 64)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ctx.ShouldBindJSON(&addHRRequest); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	addHRRequest.CompanyID = uint(cid)

	err = addHR(ctx, &addHRRequest)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	logrus.Infof("An HR %s is added with id %d", addHRRequest.Name, addHRRequest.ID)

	ctx.JSON(http.StatusOK, gin.H{"status": "Successfully added"})

}
