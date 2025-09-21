package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spo-iitk/ras-backend/constants"
	"github.com/spo-iitk/ras-backend/middleware"
)

type godLoginRequest struct {
	AdminID    string `json:"admin_id" binding:"required"`
	Password   string `json:"password" binding:"required"`
	UserID     string `json:"user_id" binding:"required"`
	RememberMe bool   `json:"remember_me"`
}

func godLoginHandler(c *gin.Context) {
	var loginReq godLoginRequest
	if err := c.ShouldBindJSON(&loginReq); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPwd, role, isActive, err := getPasswordAndRole(c, loginReq.AdminID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !comparePasswords(hashedPwd, loginReq.Password) {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid Credentials"})
		return
	}

	if role != constants.GOD && role != constants.OPC {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Only God and OPC can login"})
		return
	}

	if !isActive && role != constants.GOD {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Admin is not active"})
		return
	}

	_, role, _, err = getPasswordAndRole(c, loginReq.UserID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := middleware.GenerateToken(loginReq.UserID, uint(role), bool(loginReq.RememberMe))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"role_id": role, "user_id": loginReq.UserID, "token": token})
}
