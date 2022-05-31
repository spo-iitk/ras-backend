package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/spo-iitk/ras-backend/auth"
)

func authServer() *http.Server {
	PORT := viper.GetString("AUTH.PORT")
	r := gin.New()
	auth.Router(r)

	server := &http.Server{
		Addr:         ":" + PORT,
		Handler:      r,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
	}

	return server
}
