package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type loginRequest struct {
	UserID   string `json:"user_id" binding:"required"`
	Password string `json:"password" binding:"required"`
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

	token, err := generateToken(loginReq.UserID, role)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.SetCookie("token", token, 0, "", "", true, true)

	c.JSON(http.StatusOK, gin.H{"status": "Successfully logged in"})
}
