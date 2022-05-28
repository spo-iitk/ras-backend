package router

import (
	"github.com/gin-gonic/gin"
	"github.com/spo-iitk/ras-backend/controllers"
)

func RASRouter(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.GET("/", controllers.HelloWorldController)
	}
}
