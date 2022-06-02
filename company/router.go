package company

import (
	"github.com/gin-gonic/gin"
	"github.com/spo-iitk/ras-backend/ras"
)

func AdminRouter(r *gin.Engine) {
	admin := r.Group("/api/admin/company")
	{
		admin.GET("", ras.PlaceHolderController)      // dump all
		admin.POST("/new", ras.PlaceHolderController) // mass dump

		admin.DELETE("/:cid", ras.PlaceHolderController)
		admin.GET("/:cid", ras.PlaceHolderController)
		admin.PUT("/:cid", ras.PlaceHolderController)

		admin.GET("/:cid/hr", ras.PlaceHolderController)
		admin.POST("/:cid/hr/new", ras.PlaceHolderController)
		admin.POST("/:cid/hr/:hrid/new-auth", ras.PlaceHolderController)
		admin.PUT("/:cid/hr/:hrid", ras.PlaceHolderController)
		admin.DELETE("/:cid/hr/:hrid", ras.PlaceHolderController)

		admin.GET("/:cid/past-hires", ras.PlaceHolderController)

		admin.GET("/:cid/history", ras.PlaceHolderController)
		admin.PUT("/:cid/history/:hid", ras.PlaceHolderController)
		admin.DELETE("/:cid/history/:hid", ras.PlaceHolderController)

	}
}
