package company

import (
	"github.com/gin-gonic/gin"
	"github.com/spo-iitk/ras-backend/ras"
)

func StudentRouter(r *gin.Engine) {
	student := r.Group("/api/student")
	{
		student.POST("/create", ras.PlaceHolderController)
		student.PUT("/:id", ras.PlaceHolderController)
		student.GET("/:id", ras.PlaceHolderController)
		student.GET("/all", ras.PlaceHolderController)
		student.GET("/programs", ras.PlaceHolderController)
		student.GET("/departments", ras.PlaceHolderController)
		student.GET("/program-departments", ras.PlaceHolderController)
	}
}

func AdminRouter(r *gin.Engine) {
	admin := r.Group("/api/admin/company")
	{
		admin.GET("", ras.PlaceHolderController)      // dump all
		admin.POST("/new", ras.PlaceHolderController) // mass dump

		admin.DELETE("/:id", ras.PlaceHolderController)
		admin.GET("/:id", ras.PlaceHolderController)
		admin.PUT("/:id", ras.PlaceHolderController)

		admin.GET("/:id/hr", ras.PlaceHolderController)
		admin.POST("/:id/hr/new", ras.PlaceHolderController)
		admin.POST("/:id/hr/:id/new-auth", ras.PlaceHolderController)
		admin.PUT("/:id/hr/:id", ras.PlaceHolderController)
		admin.DELETE("/:id/hr/:id", ras.PlaceHolderController)

		admin.GET("/:id/past-hires", ras.PlaceHolderController)

		admin.GET("/:id/history", ras.PlaceHolderController)
		admin.PUT("/:id/history/:id", ras.PlaceHolderController)
		admin.DELETE("/:id/history/:id", ras.PlaceHolderController)

	}
}
