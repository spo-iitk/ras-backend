package student

import (
	"github.com/gin-gonic/gin"
	"github.com/spo-iitk/ras-backend/mail"
	"github.com/spo-iitk/ras-backend/ras"
)

func StudentRouter(r *gin.Engine) {
	student := r.Group("/api/student")
	{
		student.PUT("", updateStudentHandler)
		student.GET("", getStudentHandler)
		student.POST("/document", postStudentDocumentHandler)
		student.GET("/documents", getStudentDocumentHandler)
	}
}

func AdminRouter(mail_channel chan mail.Mail, r *gin.Engine) {
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

		admin.POST("/:sid/clarification", postClarificationHandler(mail_channel))
		admin.GET("/:sid/documents", getDocumentHandler)
		admin.PUT("/document/:docid/verify", putDocumentVerifyHandler(mail_channel))
		admin.GET("/documents", getAllDocumentHandler)
		admin.GET("/documents/type/:type", getAllDocumentHandlerByType)
	}
}
