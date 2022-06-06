package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/spo-iitk/ras-backend/application"
	"github.com/spo-iitk/ras-backend/middleware"
	"github.com/spo-iitk/ras-backend/rc"
)

func companyServer() *http.Server {
	PORT := viper.GetString("PORT.COMPANY")
	engine := gin.New()
	engine.Use(middleware.Authenticator())

	rc.CompanyRouter(engine)
	application.CompanyRouter(engine)

	server := &http.Server{
		Addr:         ":" + PORT,
		Handler:      engine,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
	}

	return server
}
