package company

import (
	"github.com/gin-gonic/gin"
	"github.com/spo-iitk/ras-backend/ras"
)

func AdminRouter(r *gin.Engine) {
	admin := r.Group("/api/admin/company")
	{
		admin.GET("", getAllCompaniesHandler)
		admin.GET("/:cid", getCompanyHandler)

		admin.PUT("", updateCompanyHandler)
		admin.POST("", addNewHandler)
		admin.POST("/bulk", addNewBulkHandler)

		admin.DELETE("/:cid", deleteCompanyHandler)

		admin.GET("/:cid/hr", getAllHRHandler)
		admin.POST("/hr", addHRHandler)
		admin.DELETE("/hr/:hrid", deleteHRHandler)

		admin.GET("/:cid/past-hires", ras.PlaceHolderController)
		admin.GET("/:cid/history", ras.PlaceHolderController)
		admin.PUT("/:cid/history/:hid", ras.PlaceHolderController)
		admin.DELETE("/:cid/history/:hid", ras.PlaceHolderController)
	}
}
