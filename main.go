package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", apiTest)
	r.Run()
	fmt.Println("API Server running ...")
}

func apiTest(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
