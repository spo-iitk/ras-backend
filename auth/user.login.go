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

	hashedPwd, role, err := getPasswordAndRole(c, loginReq.UserID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !comparePasswords(hashedPwd, loginReq.Password) {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid Credentials"})
		return
	}

	token, err := middleware.GenerateToken(loginReq.UserID, uint(role), bool(loginReq.RememberMe))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = setLastLogin(c, loginReq.UserID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.SetCookie("token", token, 0, "", "", false, false)

	c.JSON(http.StatusOK, gin.H{"status": "Successfully logged in"})
}
