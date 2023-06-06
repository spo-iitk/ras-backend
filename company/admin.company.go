package company

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	m "github.com/spo-iitk/ras-backend/middleware"
)

func addNewHandler(ctx *gin.Context) {
	var newCompanyRequest Company

	if err := ctx.ShouldBindJSON(&newCompanyRequest); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := createCompany(ctx, &newCompanyRequest)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	logrus.Infof("A new company %s is added with id %d by %s", newCompanyRequest.Name, newCompanyRequest.ID, m.GetUserID(ctx))

	ctx.JSON(http.StatusOK, gin.H{"status": "Successfully added"})

}

func addNewBulkHandler(ctx *gin.Context) {
	var request []Company

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := createCompanies(ctx, &request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	logrus.Infof("%d companies is added with by %s", len(request), m.GetUserID(ctx))

	ctx.JSON(http.StatusOK, gin.H{"status": "Successfully added"})

}

func updateCompanyHandler(ctx *gin.Context) {
	var updateCompanyRequest Company

	if err := ctx.ShouldBindJSON(&updateCompanyRequest); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if updateCompanyRequest.ID == 0 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Enter Company ID"})
		return
	}
	updated, err := updateCompany(ctx, &updateCompanyRequest)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !updated {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Company not found"})
		return
	}

	logrus.Infof("A company with id %d is updated by %s", updateCompanyRequest.ID, m.GetUserID(ctx))

	ctx.JSON(http.StatusOK, gin.H{"status": "Successfully updated"})
}

func deleteCompanyHandler(ctx *gin.Context) {

	cid, err := strconv.ParseUint(ctx.Param("cid"), 10, 32)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = deleteCompany(ctx, uint(cid))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	logrus.Infof("A company with id %d is deleted by %s", cid, m.GetUserID(ctx))

	ctx.JSON(http.StatusOK, gin.H{"status": "Successfully deleted"})

}
