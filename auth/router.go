package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/spo-iitk/ras-backend/mail"
)

func Router(mail_channel chan mail.Mail, r *gin.Engine) {
	auth := r.Group("/api/auth")
	{
		auth.POST("/login", loginHandler)
		auth.GET("/admins", getAllAdminDetailsHandler)
		auth.GET("/admins/:userID", getAdminDetailsHandler)
		auth.PUT("/admins/:userID/role", updateUserRole)
		auth.POST("/signup", signUpHandler(mail_channel))
		auth.POST("/otp", otpHandler(mail_channel))
		auth.POST("/reset-password", resetPasswordHandler(mail_channel))
		auth.POST("/company-signup", companySignUpHandler(mail_channel))

		auth.GET("/whoami", whoamiHandler) // who am i, if not exploited
		auth.GET("/credits", creditsHandler)

		auth.POST("/hr-signup", hrSignUpHandler(mail_channel))

		auth.GET("/new-companies", companiesAddedHandler)

		auth.POST("/god/signup", godSignUpHandler(mail_channel))
		auth.POST("/god/login", godLoginHandler)
		auth.POST("/god/reset-password", godResetPasswordHandler(mail_channel))
	}
}
