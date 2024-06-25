package controllers

import (
	"fmt"
	"net/http"

	"github.com/4kpros/go-crud/common/utils"
	"github.com/4kpros/go-crud/config"
	"github.com/4kpros/go-crud/services/post/models"
	"github.com/gin-gonic/gin"
)

func CreatePost(c *gin.Context) {
	// Get data of req body
	var body struct {
		Body  string
		Title string
	}
	err := c.Bind(&body)
	if err != nil {
		message := "All these fields are required: {Body: string, Title: string}"
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("%s", message))
		return
	}

	// Create a post
	post := models.Post{
		Title: body.Title,
		Body:  body.Body,
	}
	result := config.DB.Create(&post)
	if result.Error != nil {
		c.AbortWithError(http.StatusBadRequest, result.Error)
		return
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
	// Get params
	id := c.Param("id")

	// Get the post
	var post models.Post
	result := config.DB.First(&post, id)
	if result.Error != nil {
		message := "Post with id " + id + " not found !"
		c.AbortWithError(http.StatusNotFound, fmt.Errorf("%s", message))
		return
	}

	// Return the response
	c.JSON(http.StatusOK, utils.ResponseData(post))
}

func GetAllPosts(c *gin.Context) {
	// Get queries
	pagination, filters := utils.GetPaginationFiltersFromQuery(c)

	// Get the posts from DB
	var posts []models.Post
	config.DB.Scopes(utils.PaginateScope(posts, pagination, filters, config.DB)).Find(&posts)

	// Return it
	c.JSON(http.StatusOK, utils.ResponseDataWithPagination(
		posts,
		*pagination,
		*filters,
	))
}
