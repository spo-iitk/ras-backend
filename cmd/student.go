package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/spo-iitk/ras-backend/middleware"
	"github.com/spo-iitk/ras-backend/student"
)

func studentServer() *http.Server {
	PORT := viper.GetString("STUDENT.PORT")
	engine := gin.New()
	engine.Use(middleware.Authenticator())
	student.StudentRouter(engine)

	server := &http.Server{
		Addr:         ":" + PORT,
		Handler:      engine,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
	}

	return server
}
