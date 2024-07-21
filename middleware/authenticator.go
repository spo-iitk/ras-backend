package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/spo-iitk/ras-backend/constants"
)

func Authenticator() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizationHeader := ctx.GetHeader("authorization")
		if len(authorizationHeader) == 0 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized,
				gin.H{"error": "authorization header is not provided"})
			return
		}

		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized,
				gin.H{"error": "invalid authorization header format"})
			return
		}

		authorizationType := strings.ToLower(fields[0])
		if authorizationType != ("bearer") {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized,
				gin.H{"error": "bearer not found"})
			return
		}

		userID, roleID, err := validateToken(fields[1])
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized,
				gin.H{"error": "invalid token"})
			return
		}

		// cookie, err := ctx.Request.Cookie("token")
		// if err != nil {
		// 	ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		// 	return
		// }

		// userID, roleID, err := validateToken(cookie.Value)
		// if err != nil {
		// 	ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid Cookies"})
		// 	return
		// }

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

func PVFAuthenticator() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizationHeader := ctx.GetHeader("authorization")
		if len(authorizationHeader) == 0 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized,
				gin.H{"error": "authorization header is not provided"})
			return
		}

		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized,
				gin.H{"error": "invalid authorization header format"})
			return
		}

		authorizationType := strings.ToLower(fields[0])
		if authorizationType != ("bearer") {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized,
				gin.H{"error": "bearer not found"})
			return
		}

		email, pid, rid, err := validatePVFToken(fields[1])
		// ctx.JSON(http.StatusAccepted, gin.H{"email": email, "pid": pid, "rid": rid}) // to be removed
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized,
				gin.H{"error": "invalid token"})
			return
		}
		ctx.Set("email", email)
		ctx.Set("pid", pid)
		ctx.Set("rid", rid)

		ctx.Next()
	}
}

func GetEmail(ctx *gin.Context) string {
	return ctx.GetString("email")
}

func GetPVFID(ctx *gin.Context) uint {
	return ctx.GetUint("pid")
}
func GetRcID(ctx *gin.Context) uint {
	return ctx.GetUint("rid")
}
