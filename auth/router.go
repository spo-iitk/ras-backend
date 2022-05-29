package auth

import (
	"github.com/gin-gonic/gin"
)

func AuthRouter(r *gin.Engine) {
	api := r.Group("/api/login")
	{
		api.GET("/login", loginHandler)
	}
}
