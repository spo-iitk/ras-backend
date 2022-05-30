package auth

import (
	"github.com/gin-gonic/gin"
)

func AuthRouter(r *gin.Engine) {
	api := r.Group("/api/auth")
	{
		api.GET("/login", loginHandler)
		api.POST("/signup", signUpHandler)
		api.POST("/otp", otpHandler)
		api.POST("/reset-password", resetPasswordHandler)
	}
}
