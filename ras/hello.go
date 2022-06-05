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

func MailController(mailQueue chan mail.Mail) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		m := mail.Mail{
			To:      []string{viper.GetString("MAIL.WEBTEAM")},
			Subject: "It Works!",
			Body:    "Hello World!",
		}
		go func(mail mail.Mail) {
			mailQueue <- m
			log.Info("Sending mail now")
		}(m)
		c.JSON(200, gin.H{
			"message": "Mail sent",
		})
	}
	return gin.HandlerFunc(fn)
}
