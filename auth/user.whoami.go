package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spo-iitk/ras-backend/middleware"
)

func whoamiHandler(ctx *gin.Context) {
	middleware.Authenticator()(ctx)
	user_id := middleware.GetUserID(ctx)
	role_id := middleware.GetRoleID(ctx)

	if user_id == "" {
		return
	}

	var user User
	err := fetchUser(ctx, &user, user_id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"role_id": role_id, "user_id": user_id, "name": user.Name})
}
