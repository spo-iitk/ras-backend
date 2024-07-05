package ras

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spo-iitk/ras-backend/mail"
)

func HelloWorldController(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello World!",
	})
}

func PlaceHolderController(c *gin.Context) {
	c.JSON(http.StatusBadRequest, gin.H{
		"message": "Please Implement me!",
	})
}

func MailController(mail_channel chan mail.Mail) gin.HandlerFunc {
	return func(c *gin.Context) {
		mail_channel <- mail.GenerateMail("yashc22@iitk.ac.in", "Test Mail", "Hello World!")
		mail_channel <- mail.GenerateMails([]string{"yashlm1017@gmail.com", "bmerchant22@iitk.ac.in"}, "Test Mail to multiple ppl", "Hello Worlds!")
		c.JSON(http.StatusOK, gin.H{"message": "Mail sent"})
	}
}
