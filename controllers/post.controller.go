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
	var body struct {
		Body  string
		Title string
	}
	err := c.Bind(&body)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("All these fields are required: {Body: string, Title: string}"))
		return
	}

	// Create a post
	if initializers.DB == nil {
		utils.Logger.Info(
			"Database instance is nil !",
		)
		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("No database found !"))
		return
	}
	post := models.Post{
		Title: body.Title,
		Body:  body.Body,
	}
	result := initializers.DB.Create(&post)
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
	result := initializers.DB.First(&post, id)
	if result.Error != nil {
		c.AbortWithError(http.StatusNotFound, fmt.Errorf("Post with id %s not found !", id))
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
	initializers.DB.Scopes(utils.PaginateScope(posts, pagination, filters, initializers.DB)).Find(&posts)

	// Return it
	c.JSON(http.StatusOK, utils.ResponseDataWithPagination(
		posts,
		*pagination,
		*filters,
	))
}
