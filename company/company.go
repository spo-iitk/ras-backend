package company

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func addNewHandler(ctx *gin.Context) {
	var newCompanyRequest Company

	if err := ctx.ShouldBindJSON(&newCompanyRequest); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := newCompany(ctx, &newCompanyRequest)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	logrus.Infof("A new company %s is added with id %d", newCompanyRequest.Name, newCompanyRequest.ID)

	ctx.JSON(http.StatusOK, gin.H{"status": "Successfully added"})

}

func updateCompanyHandler(ctx *gin.Context) {
	var updateCompanyRequest Company

	if err := ctx.ShouldBindJSON(&updateCompanyRequest); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cid, err := strconv.ParseUint(ctx.Param("cid"), 10, 64)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updated, err := updateCompany(ctx, &updateCompanyRequest, uint(cid))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !updated {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Company not found"})
		return
	}

	logrus.Infof("A company with id %d is updated", cid)

	ctx.JSON(http.StatusOK, gin.H{"status": "Successfully updated"})
}

func deleteCompanyHandler(ctx *gin.Context) {
	var deleteCompanyRequest Company

	id, err := strconv.ParseUint(ctx.Param("hrid"), 10, 64)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = deleteCompany(ctx, &deleteCompanyRequest, uint(id))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	logrus.Infof("A company with id %d is deleted", id)

	ctx.JSON(http.StatusOK, gin.H{"status": "Successfully deleted"})

}
