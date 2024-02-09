package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/spo-iitk/ras-backend/application"
	"github.com/spo-iitk/ras-backend/company"
	"github.com/spo-iitk/ras-backend/mail"
	"github.com/spo-iitk/ras-backend/middleware"
	"github.com/spo-iitk/ras-backend/rc"
	"github.com/spo-iitk/ras-backend/student"
)

func adminRCServer(mail_channel chan mail.Mail) *http.Server {
	PORT := viper.GetString("PORT.ADMIN.RC")
	engine := gin.New()
	engine.Use(middleware.CORS())
	engine.Use(middleware.Authenticator())
	engine.Use(middleware.EnsurePsuedoAdmin())
	engine.Use(gin.CustomRecovery(recoveryHandler))
	engine.Use(gin.Logger())

	rc.AdminRouter(mail_channel, engine)

	server := &http.Server{
		Addr:         ":" + PORT,
		Handler:      engine,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
	}

	return server
}

func adminApplicationServer(mail_channel chan mail.Mail) *http.Server {
	PORT := viper.GetString("PORT.ADMIN.APP")
	engine := gin.New()
	engine.Use(middleware.CORS())
	engine.Use(middleware.Authenticator())
	engine.Use(middleware.EnsurePsuedoAdmin())
	engine.Use(gin.CustomRecovery(recoveryHandler))
	engine.Use(gin.Logger())

	application.AdminRouter(mail_channel, engine)

	server := &http.Server{
		Addr:         ":" + PORT,
		Handler:      engine,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
	}

	return server
}

func adminCompanyServer() *http.Server {
	PORT := viper.GetString("PORT.ADMIN.COMPANY")
	engine := gin.New()
	engine.Use(middleware.CORS())
	engine.Use(middleware.Authenticator())
	engine.Use(middleware.EnsureAdmin())
	engine.Use(gin.CustomRecovery(recoveryHandler))
	engine.Use(gin.Logger())

	company.AdminRouter(engine)

	server := &http.Server{
		Addr:         ":" + PORT,
		Handler:      engine,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
	}

	return server
}

func adminStudentServer() *http.Server {
	PORT := viper.GetString("PORT.ADMIN.STUDENT")
	engine := gin.New()
	engine.Use(middleware.CORS())
	engine.Use(middleware.Authenticator())
	engine.Use(middleware.EnsurePsuedoAdmin())
	engine.Use(gin.CustomRecovery(recoveryHandler))

	student.AdminRouter(engine)

	server := &http.Server{
		Addr:         ":" + PORT,
		Handler:      engine,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
	}

	return server
}
