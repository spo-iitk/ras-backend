package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/spo-iitk/ras-backend/auth"
	"github.com/spo-iitk/ras-backend/mail"
	"github.com/spo-iitk/ras-backend/middleware"
)

func authServer(mail_channel chan mail.Mail) *http.Server {
	PORT := viper.GetString("PORT.AUTH")
	r := gin.New()
	r.Use(middleware.CORS())
	r.Use(gin.CustomRecovery(recoveryHandler))
	r.Use(gin.Logger())

	auth.Router(mail_channel, r)

	server := &http.Server{
		Addr:         ":" + PORT,
		Handler:      r,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
	}

	return server
}
