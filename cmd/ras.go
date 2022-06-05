package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spo-iitk/ras-backend/mail"
	"github.com/spo-iitk/ras-backend/ras"
)

func rasServer(mail_channel chan mail.Mail) *http.Server {
	PORT := "8080"
	engine := gin.New()
	// engine.Use(middleware.Authenticator())
	ras.RASRouter(mail_channel, engine)

	server := &http.Server{
		Addr:         ":" + PORT,
		Handler:      engine,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
	}
	return server
}
