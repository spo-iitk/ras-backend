package company

import (
	"github.com/gin-gonic/gin"
	"github.com/spo-iitk/ras-backend/middleware"
)

func extractCompanyID(ctx *gin.Context) (uint, error) {
	user_email := middleware.GetUserID(ctx)
	return FetchCompanyIDByEmail(ctx, user_email)
}
