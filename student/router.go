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
		admin.GET("", getAllStudentsHandler)
		admin.GET("/:sid", getStudentByIDHandler)
		admin.PUT("/:sid", updateStudentByIDHandler)
		admin.GET("/:sid/history", ras.PlaceHolderController)
	}
}
