package ras

import "github.com/gin-gonic/gin"

func HelloWorldController(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello World!",
	})
}

func PlaceHolderController(c *gin.Context) {
	c.JSON(500, gin.H{
		"message": "Please Implement me!",
	})
}
