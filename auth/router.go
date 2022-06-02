package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/spo-iitk/ras-backend/ras"
)

func Router(r *gin.Engine) {
	auth := r.Group("/api/auth")
	{
		auth.POST("/login", loginHandler)
		auth.POST("/signup", signUpHandler)
		auth.POST("/otp", otpHandler)
		auth.POST("/reset-password", resetPasswordHandler)
		auth.POST("/company-signup", companySignUpHandler)
		auth.GET("/whoami", ras.PlaceHolderController) // who am i, if not exploited
	}
}
