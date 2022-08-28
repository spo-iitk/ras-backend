package company

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spo-iitk/ras-backend/middleware"
)

func postNewHRHandler(ctx *gin.Context) {
	var addHRRequest CompanyHR

	if err := ctx.ShouldBindJSON(&addHRRequest); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if addHRRequest.CompanyID != 0 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Company ID is not allowed"})
		return
	}

	// companyID, err := rc.ExtractCompanyID(ctx)
	user_email := middleware.GetUserID(ctx)
	companyID, err := FetchCompanyIDByEmail(ctx, user_email)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	addHRRequest.CompanyID = companyID

	err = addHR(ctx, &addHRRequest)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "Successfully added"})

}

type updateHRRequest struct {
	Name        string `json:"name"`
	Phone       string `json:"phone"`
	Designation string `json:"designation"`
}

func putHRHandler(ctx *gin.Context) {
	var req updateHRRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hrid := middleware.GetUserID(ctx)
	companyID, err := FetchCompanyIDByEmail(ctx, hrid)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = updateHR(ctx, companyID, hrid, &req)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "Successfully updated"})
}
