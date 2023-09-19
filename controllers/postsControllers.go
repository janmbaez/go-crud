package controllers

import (
	"go-crud/initializers"
	"go-crud/models"

	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {
	// Get data of req body
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)
	// Create a post

	post := models.Post{Title: body.Title, Body: body.Body}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsIndex(c *gin.Context) {
	// Get the posts
	var posts []models.Post
	initializers.DB.Find(&posts)

	c.JSON(200, gin.H{
		"posts": posts,
	})
	// Respond with them
}

func PostsShow(c *gin.Context) {
	// Get id off url
	id := c.Param("id")

	// Get the posts
	var post models.Post
	initializers.DB.First(&post, id)

	// Respond with them
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsUpdate(c *gin.Context) {
	// Get the id off the url
	id := c.Param("id")
	// Get the data off rq body
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)
	// Find the post were updating
	var post models.Post
	initializers.DB.First(&post, id)
	// Update
	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title, Body: body.Body,
	})

	// Respond with it
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsDelete(c *gin.Context) {
	// Get the id off the url
	id := c.Param("id")
	// Delete the posts
	initializers.DB.Delete(&models.Post{}, id)
	// Respond
	c.Status(200)
}
