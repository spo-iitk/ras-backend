package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spo-iitk/ras-backend/middleware"
)

func whoamiHandler(ctx *gin.Context) {
	middleware.Authenticator()(ctx)
	data := gin.H{"role": middleware.GetRoleID(ctx), "user_id": middleware.GetUserID(ctx)}
	ctx.JSON(http.StatusOK, gin.H{"status": "User Logged In", "data": data})
}
