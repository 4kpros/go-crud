package main

import (
	"github.com/4kpros/go-crud/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVars()
}

func main() {
	r := gin.Default()
	r.GET("/", apiTest)
	r.Run()
}

func apiTest(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Welcome to GO-CRUD API",
	})
}
