package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spo-iitk/ras-backend/constants"
	"github.com/spo-iitk/ras-backend/middleware"
)

func companiesAddedHandler(ctx *gin.Context) {
	middleware.Authenticator()(ctx)
	role := middleware.GetRoleID(ctx)
	if role != constants.OPC && role != constants.GOD {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	companies, err := getAllCompaniesAdded(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"companies": companies})
}
