package auth

import (
	"github.com/gin-gonic/gin"
)

func AuthRouter(r *gin.Engine) {
	api := r.Group("/api/auth")
	{
		api.GET("/login", loginHandler)
	}
}
//login signup forgotpass otp companysignup
//hashing salting
//jwt token