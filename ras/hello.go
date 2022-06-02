package ras

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

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
		m := mail.Mail{
			Sender:  viper.GetString("MAIL.USER") + "@iitk.ac.in",
			To:      viper.GetStringSlice("MAIL.WEBTEAM"),
			Bcc:     []string{"shreea20@iitk.ac.in"},
			Subject: "Hi vro",
			Body:    "Test",
		}
		go func(mail mail.Mail) {
			mail_channel <- m
			log.Info("Sending mail now")
		}(m)
		c.JSON(200, gin.H{
			"message": "Mail sent",
		})
	}
	return gin.HandlerFunc(fn)
}
