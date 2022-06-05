package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spo-iitk/ras-backend/mail"
)

type signUpRequest struct {
	UserID   string `json:"user_id" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
	OTP      string `json:"otp" binding:"required"`
}

func signUpHandler(mail_channel chan mail.Mail) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var signupReq signUpRequest

		if err := ctx.ShouldBindJSON(&signupReq); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		verified, err := verifyOTP(ctx, signupReq.UserID, signupReq.OTP)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if !verified {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid OTP"})
			return
		}

		hashedPwd := hashAndSalt(signupReq.Password)

		id, err := createUser(ctx, &User{
			UserID:   signupReq.UserID,
			Name:     signupReq.Name,
			Password: hashedPwd,
		})

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		logrus.Infof("User %s created successfully with id %d", signupReq.UserID, id)
		mail_channel <- mail.GenerateMail(signupReq.UserID, signupReq.Name+"registered on RAS", "Dear "+signupReq.Name+",\n\nYou have been registered on RAS")
		ctx.JSON(http.StatusOK, gin.H{"status": "Successfully signed up"})
	}
}
