package middleware

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/spo-iitk/ras-backend/constants"
)

func EnsureAdmin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		role, err := strconv.ParseUint(GetRoleID(ctx), 10, 64)
		if err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		roleID := constants.Role(role)

		if roleID != constants.OPC && roleID != constants.GOD {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		ctx.Next()
	}
}
