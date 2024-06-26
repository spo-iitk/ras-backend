package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/spo-iitk/ras-backend/mail"
	"github.com/spo-iitk/ras-backend/middleware"
	"github.com/spo-iitk/ras-backend/ras"
)

func rasServer(mail_channel chan mail.Mail) *http.Server {
	PORT := viper.GetString("PORT.RAS")
	engine := gin.New()
	engine.Use(middleware.CORS())
	// engine.Use(middleware.Authenticator())
	ras.RASRouter(mail_channel, engine)
	engine.Use(gin.CustomRecovery(recoveryHandler))
	engine.Use(gin.Logger())

	server := &http.Server{
		Addr:         ":" + PORT,
		Handler:      engine,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
	}
	return server
}
