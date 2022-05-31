package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
	"github.com/spo-iitk/ras-backend/auth"
)

var signingKey = []byte(viper.GetString("JWT.PRIVATE_KEY"))

func Authenticator() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		cookie, err := ctx.Request.Cookie("token")
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		userID, roleID, err := validateToken(cookie.Value)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}

		ctx.Set("userID", userID)
		ctx.Set("roleID", roleID)

		ctx.Next()
	}
}

func validateToken(encodedToken string) (string, auth.Role, error) {

	claims := &auth.CustomClaims{}
	_, err := jwt.ParseWithClaims(encodedToken, claims, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, fmt.Errorf("invalid token %s", token.Header["alg"])
		}
		return []byte(signingKey), nil
	})

	if err != nil {
		return "", 12, nil
	}
	return claims.UserID, claims.RoleID, nil
}

func GetUserID(ctx *gin.Context) string {
	return ctx.GetString("userID")
}

func GetRoleID(ctx *gin.Context) string {
	return ctx.GetString("roleID")
}
