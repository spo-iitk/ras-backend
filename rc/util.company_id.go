package rc

import (
	"github.com/gin-gonic/gin"
	"github.com/spo-iitk/ras-backend/company"
	"github.com/spo-iitk/ras-backend/middleware"
)

// func extractCompanyRCID(ctx *gin.Context) (uint, error) {
// 	companyID, err := extractCompanyID(ctx)
// 	if err != nil {
// 		return 0, err
// 	}

// 	rid, err := util.ParseUint(ctx.Param("rid"))
// 	if err != nil {
// 		return 0, err
// 	}

// 	return FetchCompanyRCID(ctx, rid, companyID)
// }

func extractCompanyID(ctx *gin.Context) (uint, error) {
	user_email := middleware.GetUserID(ctx)
	return company.FetchCompanyIDByEmail(ctx, user_email)
}
