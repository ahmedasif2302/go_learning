package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "go lang init api call",
		})
	})
	fmt.Print("Server started at the port: 2000")
	router.Run(":2000")
}
