package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spo-iitk/ras-backend/constants"
	"github.com/spo-iitk/ras-backend/mail"
	"github.com/spo-iitk/ras-backend/middleware"
)

type godResetPasswordRequest struct {
	UserID      string `json:"user_id" binding:"required"`
	NewPassword string `json:"new_password" binding:"required"`
}

func godResetPasswordHandler(mail_channel chan mail.Mail) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		middleware.Authenticator()(ctx)
		if middleware.GetRoleID(ctx) != constants.GOD && middleware.GetRoleID(ctx) != constants.OPC {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Only OPC and GOD can access"})
			return
		}

		var resetPasswordReq godResetPasswordRequest

		if err := ctx.ShouldBindJSON(&resetPasswordReq); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		hashedPwd := hashAndSalt(resetPasswordReq.NewPassword)

		ok, err := updatePasswordbyGod(ctx, resetPasswordReq.UserID, hashedPwd)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if !ok {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "No such student exists"})
			return
		}

		logrus.Infof("Password of %s reset successfully", resetPasswordReq.UserID)
		mail_channel <- mail.GenerateMail(resetPasswordReq.UserID, "Password Reset Successfully", "Your password has been reset successfully. Your new password is: "+resetPasswordReq.NewPassword)

		ctx.JSON(http.StatusOK, gin.H{"status": "Successfully reset password"})
	}
}
