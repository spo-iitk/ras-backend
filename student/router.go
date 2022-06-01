package student

import (
	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {
	student := r.Group("/api/student")
	{
		student.POST("/create", createStudentHandler)
		student.PUT("/:id", updateStudentHandler)
		student.GET("/:id", getStudentHandler)
		student.GET("/all", getAllStudentsHandler)
		student.GET("/programs", getPrograms)
		student.GET("/departments", getDepartments)
		student.GET("/program-departments", getProgramsDepartments)
	}
}
