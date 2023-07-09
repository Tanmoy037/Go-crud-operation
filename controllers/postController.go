package controllers

import (
	"github.com/Tanmoy037/go-crud/initializers"
	"github.com/Tanmoy037/go-crud/models"
	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {
	//get data off req body
	var body struct {
		Body  string `json:"body"`
		Title string `json:"title"`
	}

	c.Bind(&body)

	// create a post
	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}
	// Return it
	c.JSON(200, gin.H{
		"post": post,
	})
}
func PostsIndex(c *gin.Context) {
	// get the posts
	var posts []models.Post
	initializers.DB.Find(&posts)

	//responed with them
	c.JSON(200, gin.H{
		"posts": posts,
	})
}
func PostsShow(c *gin.Context) {
	// get if off url
	id := c.Param("id")

	// get the posts
	var post models.Post
	initializers.DB.First(&post, id)

	// Respond with them
	c.JSON(200, gin.H{
		"post": post,
	})
}
func PostUpdate(c *gin.Context) {
	// get the id off the url
	id := c.Param("id")

	//get the data off req body
	var body struct {
		Body  string `json:"body"`
		Title string `json:"title"`
	}

	c.Bind(&body)

	//Find the post were updating
	var post models.Post
	initializers.DB.First(&post, id)

	//update it
	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})


	//respond with it
	c.JSON(200, gin.H{
		"post": post,
	})

}

func PostDelete(c *gin.Context) {
	//get the id off the url
	id := c.Param("id")

	//delete the post 
	initializers.DB.Delete(&models.Post{}, id)


	// Respond with it
	c.Status(200)
}
