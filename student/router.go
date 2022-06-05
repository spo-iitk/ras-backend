package student

import (
	"github.com/gin-gonic/gin"
	"github.com/spo-iitk/ras-backend/ras"
)

func StudentRouter(r *gin.Engine) {
	student := r.Group("/api/student")
	{
		student.POST("/new", createStudentHandler)
		student.PUT("/:sid", updateStudentHandler)
		student.GET("/:sid", getStudentHandler)
		student.GET("/all", getAllStudentsHandler)
	}
}

func AdminRouter(r *gin.Engine) {
	admin := r.Group("/api/admin/student")
	{
		admin.GET("", ras.PlaceHolderController)              // dump all
		admin.GET("/:sid", ras.PlaceHolderController)         // dump all
		admin.PUT("/:sid", ras.PlaceHolderController)         // dump all
		admin.GET("/:sid/history", ras.PlaceHolderController) // mass dump
	}
}
