package controllers

import (
	"fmt"
	"net/http"

	"github.com/4kpros/go-crud/initializers"
	"github.com/4kpros/go-crud/models"
	"github.com/4kpros/go-crud/utils"
	"github.com/gin-gonic/gin"
)

func CreatePost(c *gin.Context) {
	// Get data of req body

	// Create a post
	post := models.Post{
		Title: "Hello",
		Body:  "Fist post",
	}
	if initializers.DB == nil {
		utils.Logger.Info(
			"database instance is nil",
		)
		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("no database found"))
		return
	}
	result := initializers.DB.Create(&post)
	if result.Error != nil {
		c.AbortWithError(http.StatusBadRequest, result.Error)
	}

	// Return it
	c.JSON(http.StatusOK, utils.ResponseData(post))
}

func UpdatePost(c *gin.Context) {
	c.JSON(http.StatusOK, utils.ResponseData(nil))
}

func DeletePost(c *gin.Context) {
	c.JSON(http.StatusOK, utils.ResponseData(nil))
}

func GetPost(c *gin.Context) {
	c.JSON(http.StatusOK, utils.ResponseData(nil))
}

func GetAllPosts(c *gin.Context) {
	c.JSON(http.StatusOK, utils.ResponseData(nil))
}
