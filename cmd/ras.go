package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spo-iitk/ras-backend/middleware"
	"github.com/spo-iitk/ras-backend/router"
)

func rasServer() *http.Server {
	PORT := "8080"
	engine := gin.New()
	engine.Use(middleware.Authenticator())
	router.RASRouter(engine)

	server := &http.Server{
		Addr:         ":" + PORT,
		Handler:      engine,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
	}
	return server
}
