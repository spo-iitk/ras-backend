package rc

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spo-iitk/ras-backend/company"
	"github.com/spo-iitk/ras-backend/middleware"
)

type companyWhoamiResponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func companyWhoamiHandler(ctx *gin.Context) {
	companyID, err := extractCompanyID(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	name, err := company.GetCompanyName(ctx, companyID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, companyWhoamiResponse{
		Name:  name,
		Email: middleware.GetUserID(ctx),
	})
}
