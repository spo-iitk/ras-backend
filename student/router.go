package student

import (
	"github.com/gin-gonic/gin"
	"github.com/spo-iitk/ras-backend/ras"
)

func StudentRouter(r *gin.Engine) {
	student := r.Group("/api/student")
	{
		student.PUT("", updateStudentHandler)
		student.GET("", getStudentHandler)
	}
}

func AdminRouter(r *gin.Engine) {
	admin := r.Group("/api/admin/student")
	{
		admin.DELETE("/:sid", deleteStudentHandler)
		admin.GET("", getAllStudentsHandler)
		admin.GET("/limited", getLimitedStudentsHandler)
		admin.PUT("", updateStudentByIDHandler)
		admin.GET("/:sid", getStudentByIDHandler)
		admin.PUT("/:sid/editable",makeStudentEdiatableHandler)
		admin.PUT("/:sid/verify", verifyStudentHandler)
		admin.GET("/:sid/history", ras.PlaceHolderController)
	}
}
