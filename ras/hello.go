package ras

import (
	"github.com/gin-gonic/gin"
	"github.com/spo-iitk/ras-backend/mail"
)

func HelloWorldController(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello World!",
	})
}

func PlaceHolderController(c *gin.Context) {
	c.JSON(500, gin.H{
		"message": "Please Implement me!",
	})
}

func MailController(mail_channel chan mail.Mail) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		mail_channel <- mail.GenerateMail("harshitr20@iitk.ac.in", "Test Mail", "Hello World!")
		mail_channel <- mail.GenerateMails([]string{"shreea20@iitk.ac.in", "ias@iitk.ac.in"}, "Test Mail to multiple ppl", "Hello Worlds!")
		c.JSON(200, gin.H{
			"message": "Mail sent",
		})
	}
	return gin.HandlerFunc(fn)
}
