package middleware

import (
	"net/http"
	"log"
	
	"github.com/gin-gonic/gin"
	"github.com/spo-iitk/ras-backend/constants"
)

func EnsureAdmin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		role := GetRoleID(ctx)

		if role != constants.OPC && role != constants.GOD {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		ctx.Next()
	}
}

func EnsurePsuedoAdmin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		role := GetRoleID(ctx)

		if role != constants.OPC && role != constants.GOD && role != constants.APC {
			log.Println(role)
			log.Println(GetUserID(ctx))
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		ctx.Next()
	}
}
