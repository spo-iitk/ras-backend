package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spo-iitk/ras-backend/middleware"
)

type loginRequest struct {
	UserID     string `json:"user_id" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RememberMe bool   `json:"remember_me"`
}

func loginHandler(c *gin.Context) {
	var loginReq loginRequest
	if err := c.ShouldBindJSON(&loginReq); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPwd, role, isActive, err := getPasswordAndRole(c, loginReq.UserID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !comparePasswords(hashedPwd, loginReq.Password) {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid Credentials"})
		return
	}

	if !isActive {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "User is not active"})
		return
	}

	token, err := middleware.GenerateToken(loginReq.UserID, uint(role), bool(loginReq.RememberMe))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	go setLastLogin(loginReq.UserID)

	c.JSON(http.StatusOK, gin.H{"role_id": role, "user_id": loginReq.UserID, "token": token})
}
