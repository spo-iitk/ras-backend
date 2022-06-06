package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/spo-iitk/ras-backend/application"
	"github.com/spo-iitk/ras-backend/company"
	"github.com/spo-iitk/ras-backend/middleware"
	"github.com/spo-iitk/ras-backend/rc"
	"github.com/spo-iitk/ras-backend/student"
)

func adminRCServer() *http.Server {
	PORT := viper.GetString("PORT.ADMIN.RC")
	engine := gin.New()
	engine.Use(middleware.Authenticator())

	rc.AdminRouter(engine)
	application.AdminRouter(engine)

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
	engine.Use(middleware.Authenticator())

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
	engine.Use(middleware.Authenticator())

	student.AdminRouter(engine)

	server := &http.Server{
		Addr:         ":" + PORT,
		Handler:      engine,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
	}

	return server
}
