package auth

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spo-iitk/ras-backend/constants"
	"github.com/spo-iitk/ras-backend/middleware"
)

type UserDetails struct {
	UserID       uint           `json:"user_id" binding:"required"`
	Password     string         `json:"password" binding:"required"`
	RoleID       constants.Role `json:"role_id" binding:"required"` // student role by default
	Name         string         `json:"name" binding:"required"`
	IsActive     bool           `json:"is_active" binding:"required"`
	LastLogin    uint           `json:"last_login" binding:"required"`
	RefreshToken string         `json:"refresh_token" binding:"required"`
}

type UpdateRoleRequest struct {
	UserID    uint           `json:"user_id" binding:"required"`
	NewRoleID constants.Role `json:"new_role_id" binding:"required"`
}

func getAllAdminDetailsHandler(ctx *gin.Context) {
	var users []User

	middleware.Authenticator()(ctx)
	middleware.EnsurePsuedoAdmin()(ctx)
	if middleware.GetUserID(ctx) == "" {
		return
	}

	if middleware.GetRoleID(ctx) < 100 {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Only admin can access this page"})
		return
	}
	err := fetchAllAdminDetails(ctx, &users)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	ctx.JSON(http.StatusOK, gin.H{"users": users})
}
func getAdminDetailsHandler(ctx *gin.Context) {
	var user User

	middleware.Authenticator()(ctx)
	middleware.EnsurePsuedoAdmin()(ctx)
	if middleware.GetUserID(ctx) == "" {
		return
	}

	err := fetchAdminDetailsById(ctx, &user, ctx.Param("userID"))

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
	currentRoleID, err := getUserRole(ctx, updateReq.UserID)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	middleware.Authenticator()(ctx)
	middleware.EnsurePsuedoAdmin()(ctx)
	if middleware.GetUserID(ctx) == "" {
		return
	}
	var userId = middleware.GetUserID(ctx)

	_, userRole, _, err := getPasswordAndRole(ctx, userId)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	if userRole > currentRoleID || userRole > updateReq.NewRoleID || userRole > 101 {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized to update this user's role"})
		return
	}

	err = updateRoleByAdmin(ctx, updateReq.UserID, updateReq.NewRoleID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadGateway, gin.H{"error": err.Error()})
	}

	logrus.Infof("User %v role changed from %v to %v - Action taken by user with id %v", updateReq.UserID, currentRoleID, updateReq.NewRoleID, userId)
	ctx.JSON(http.StatusOK, gin.H{"message": "User role updated successfully"})
}

func updateUserActiveStatus(ctx *gin.Context) {
	requestedUserId, err := strconv.ParseUint(ctx.Param("userID"), 10, 16)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	middleware.Authenticator()(ctx)
	middleware.EnsurePsuedoAdmin()(ctx)

	userId := middleware.GetUserID(ctx)
	roleId := middleware.GetRoleID(ctx)

	var requestedUserRoleID constants.Role
	requestedUserRoleID, err = getUserRole(ctx, uint(requestedUserId))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if roleId > requestedUserRoleID && roleId > 101 {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized to update this user's activity status"})
		return
	}

	active, err := toggleActive(ctx, uint(requestedUserId))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}

	logrus.Infof("User %v active status set to %v - Action taken by user with id %v", requestedUserId, active, userId)
	ctx.JSON(http.StatusOK, gin.H{"message": "User status updated successfully"})

}
