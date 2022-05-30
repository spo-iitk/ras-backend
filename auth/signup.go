package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type signUpRequest struct {
	UserID   string `json:"user_id" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
	OTP      string `json:"otp" binding:"required"`
}

func signUpHandler(ctx *gin.Context) {
	var signupReq signUpRequest

	if err := ctx.ShouldBindJSON(&signupReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	verified, err := verifyOTP(ctx, signupReq.UserID, signupReq.OTP)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !verified {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid OTP"})
		return
	}

	hashedPwd := hashAndSalt(signupReq.Password)

	id, err := createUser(ctx, &User{
		UserID:   signupReq.UserID,
		Name:     signupReq.Name,
		Password: hashedPwd,
	})

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	logrus.Infof("User %s created successfully with id %d", signupReq.UserID, id)

	ctx.JSON(http.StatusOK, gin.H{"status": "Successfully signed up"})
}
