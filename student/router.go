package student

import (
	"github.com/gin-gonic/gin"
	"github.com/spo-iitk/ras-backend/ras"
)

func StudentRouter(r *gin.Engine) {
	student := r.Group("/api/student")
	{
		student.PUT("/:sid", updateStudentHandler)
		student.GET("/:sid", getStudentHandler)
	}
}

func AdminRouter(r *gin.Engine) {
	admin := r.Group("/api/admin/student")
	{
		admin.GET("", getAllStudentsHandler)              
		admin.GET("/:sid", getStudentHandler)         
		admin.PUT("/:sid", updateStudentHandler)        
		admin.GET("/:sid/history", ras.PlaceHolderController) 
	}
}
