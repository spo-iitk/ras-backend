package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spo-iitk/ras-backend/constants"
	"github.com/spo-iitk/ras-backend/middleware"
)

type UserDetails struct {
	UserID       string         `json:"user_id" binding:"required"`
	Password     string         `json:"password" binding:"required"`
	RoleID       constants.Role `json:"role_id" binding:"required"` // student role by default
	Name         string         `json:"name" binding:"required"`
	IsActive     bool           `json:"is_active" binding:"required"`
	LastLogin    uint           `json:"last_login" binding:"required"`
	RefreshToken string         `json:"refresh_token" binding:"required"`
}

type UpdateRoleRequest struct {
	UserID    string         `json:"user_id" binding:"required"`
	NewRoleID constants.Role `json:"new_role_id" binding:"required"`
}

func getAllAdminDetailsHandler(ctx *gin.Context) {
	var users []User

	middleware.Authenticator()(ctx)
	if middleware.GetUserID(ctx) == "" {
		return
	}

	if middleware.GetRoleID(ctx) < 100 {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Only admin can access this page"})
		return
	}
	err := fetchAdmins(ctx, &users)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	ctx.JSON(http.StatusOK, gin.H{"users": users})
}
func getAdminDetailsHandler(ctx *gin.Context) {
	var user User

	middleware.Authenticator()(ctx)
	if middleware.GetUserID(ctx) == "" {
		return
	}

	err := fetchAdmin(ctx, &user, ctx.Param("userID"))

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	ctx.JSON(http.StatusOK, user)
}
func updateUserRole(ctx *gin.Context) {

	var updateReq UpdateRoleRequest

	if err := ctx.ShouldBindJSON(&updateReq); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var currentRoleID constants.Role
	_, currentRoleID, _, err := getPasswordAndRole(ctx, updateReq.UserID)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	middleware.Authenticator()(ctx)
	if middleware.GetUserID(ctx) == "" {
		return
	}

	if middleware.GetRoleID(ctx) > currentRoleID {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized to update this user's role"})
		return
	}

	err = updateRole(ctx, updateReq.UserID, updateReq.NewRoleID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadGateway, gin.H{"error": err.Error()})
	}

	logrus.New().Infof("User %d role changed from %d to %d", updateReq.UserID, currentRoleID, updateReq.NewRoleID)
	ctx.JSON(http.StatusOK, gin.H{"message": "User role updated successfully"})
}

// Active inactive ka dekhna hai
