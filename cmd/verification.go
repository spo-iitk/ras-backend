package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/spo-iitk/ras-backend/application"
	"github.com/spo-iitk/ras-backend/mail"
	"github.com/spo-iitk/ras-backend/middleware"
)

func verificationServer(mail_channel chan mail.Mail) *http.Server {
	PORT := viper.GetString("PORT.VERIFICATION")
	fmt.Print(PORT)
	engine := gin.New()
	engine.Use(middleware.CORS())
	engine.Use(middleware.PVFAuthenticator())
	engine.Use(gin.CustomRecovery(recoveryHandler))
	engine.Use(gin.Logger())

	application.PvfVerificationRouter(mail_channel, engine)

	server := &http.Server{
		Addr:         ":" + PORT,
		Handler:      engine,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
	}

	return server
}
