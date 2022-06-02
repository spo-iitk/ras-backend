package ras

import (
	"github.com/gin-gonic/gin"
)

func RASRouter(r *gin.Engine) {
	api := r.Group("/api/ras")
	{
		api.GET("", HelloWorldController)
		api.GET("/programs", PlaceHolderController)
		api.GET("/departments", PlaceHolderController)
		api.GET("/program-departments", PlaceHolderController)
	}
}
