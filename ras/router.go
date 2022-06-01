package ras

import (
	"github.com/gin-gonic/gin"
)

func RASRouter(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.GET("/", HelloWorldController)
	}
}
