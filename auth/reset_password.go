package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type resetPasswordRequest struct {
	UserID      string `json:"user_id" binding:"required"`
	NewPassword string `json:"New_password" binding:"required"`
	OTP         string `json:"otp" binding:"required"`
}

func resetPasswordHandler(ctx *gin.Context) {
	var resetPasswordReq resetPasswordRequest

	if err := ctx.ShouldBindJSON(&resetPasswordReq); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	verified, err := verifyOTP(ctx, resetPasswordReq.UserID, resetPasswordReq.OTP)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !verified {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid OTP"})
		return
	}

	hashedPwd := hashAndSalt(resetPasswordReq.NewPassword)

	err = updatePassword(ctx, resetPasswordReq.UserID, hashedPwd)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	logrus.Infof("Password of %s reset successfully", resetPasswordReq.UserID)

	ctx.JSON(http.StatusOK, gin.H{"status": "Successfully reset password up"})
}
