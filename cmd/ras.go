package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spo-iitk/ras-backend/router"
)

func rasRouter() *http.Server {
	PORT := "8080"
	r := gin.New()
	router.RASRouter(r)

	server := &http.Server{
		Addr:         ":" + PORT,
		Handler:      r,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
	}
	return server
}
