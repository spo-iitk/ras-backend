package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/spo-iitk/ras-backend/application"
	"github.com/spo-iitk/ras-backend/mail"
	"github.com/spo-iitk/ras-backend/middleware"
	"github.com/spo-iitk/ras-backend/rc"
	"github.com/spo-iitk/ras-backend/student"
)

func studentServer(mail_channel chan mail.Mail) *http.Server {
	PORT := viper.GetString("PORT.STUDENT")
	engine := gin.New()
	engine.Use(middleware.CORS())
	engine.Use(middleware.Authenticator())
	engine.Use(gin.CustomRecovery(recoveryHandler))
	engine.Use(gin.Logger())

	student.StudentRouter(engine)
	rc.StudentRouter(engine)
	application.StudentRouter(mail_channel, engine)

	server := &http.Server{
		Addr:         ":" + PORT,
		Handler:      engine,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
	}

	return server
}
