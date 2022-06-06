package middleware

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/spo-iitk/ras-backend/constants"
)

func Authenticator() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		cookie, err := ctx.Request.Cookie("token")
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		userID, roleID, err := validateToken(cookie.Value)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid Cookies"})
			return
		}

		ctx.Set("userID", userID)
		ctx.Set("roleID", roleID)

		ctx.Next()
	}
}

func GetUserID(ctx *gin.Context) string {
	return ctx.GetString("userID")
}

func GetRoleID(ctx *gin.Context) constants.Role {
	role, err := strconv.ParseUint(ctx.GetString("roleID"), 10, 8)
	if err != nil {
		return constants.NONE
	}

	return constants.Role(role)
}
