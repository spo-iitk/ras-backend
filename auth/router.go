package auth

import (
	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {
	auth := r.Group("/api/auth")
	{
		auth.POST("/login", loginHandler)
		auth.POST("/signup", signUpHandler)
		auth.POST("/otp", otpHandler)
		auth.POST("/reset-password", resetPasswordHandler)
		auth.POST("/company-signup", companySignUpHandler)
	}
}
