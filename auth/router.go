package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/spo-iitk/ras-backend/mail"
	"github.com/spo-iitk/ras-backend/ras"
)

func Router(mail_channel chan mail.Mail, r *gin.Engine) {
	auth := r.Group("/api/auth")
	{
		auth.POST("/login", loginHandler)
		auth.POST("/signup", signUpHandler(mail_channel))
		auth.POST("/otp", otpHandler(mail_channel))
		auth.POST("/reset-password", resetPasswordHandler(mail_channel))
		auth.POST("/company-signup", companySignUpHandler(mail_channel))
		auth.GET("/whoami", ras.PlaceHolderController) // who am i, if not exploited
	}
}
