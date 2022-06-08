package middleware

import (
	"net/http"

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
		ctx.Set("roleID", int(roleID))

		ctx.Next()
	}
}

func GetUserID(ctx *gin.Context) string {
	return ctx.GetString("userID")
}

func GetRoleID(ctx *gin.Context) constants.Role {

	return constants.Role(ctx.GetInt("roleID"))
}
