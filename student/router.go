package student

import (
	"github.com/gin-gonic/gin"
	"github.com/spo-iitk/ras-backend/ras"
)

func StudentRouter(r *gin.Engine) {
	student := r.Group("/api/student")
	{
		student.POST("/new", createStudentHandler)
		student.PUT("/:id", updateStudentHandler)
		student.GET("/:id", getStudentHandler)
		student.GET("/all", getAllStudentsHandler)
		student.GET("/programs", getPrograms)
		student.GET("/departments", getDepartments)
		student.GET("/program-departments", getProgramsDepartments)
	}
}

func AdminRouter(r *gin.Engine) {
	student := r.Group("/api/admin/student")
	{
		student.GET("", ras.PlaceHolderController)             // dump all
		student.GET("/:id", ras.PlaceHolderController)         // dump all
		student.PUT("/:id", ras.PlaceHolderController)         // dump all
		student.GET("/:id/history", ras.PlaceHolderController) // mass dump
	}
}
