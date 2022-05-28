package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/spo-iitk/ras-backend/router"
)

func main() {
	r := gin.Default()
	router.RASRouter(r)

	if err := r.Run(); err != nil {
		log.Fatalln("[ERROR] Could not start the server: ", err.Error())
	}
}
