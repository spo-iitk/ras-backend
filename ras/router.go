package ras

import (
	"github.com/gin-gonic/gin"
	"github.com/spo-iitk/ras-backend/mail"
)

func RASRouter(mailQueue chan mail.Mail, r *gin.Engine) {
	api := r.Group("/api/ras")
	{
		api.GET("", HelloWorldController)
		api.GET("/programs", PlaceHolderController)
		api.GET("/departments", PlaceHolderController)
		api.GET("/program-departments", PlaceHolderController)
		api.GET("/testmail", MailController(mailQueue))
	}
}
