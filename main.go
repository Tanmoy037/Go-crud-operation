package main

import (
	"github.com/Tanmoy037/go-crud/controllers"
	"github.com/Tanmoy037/go-crud/initializers"

	//	"github.com/Tanmoy037/go-crud/models"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()
	r.POST("/posts", controllers.PostsCreate)
	r.PUT("/posts/:id", controllers.PostUpdate)
	r.GET("/posts", controllers.PostsIndex)
	r.GET("/posts/:id", controllers.PostsShow)
	r.DELETE("/posts/:id", controllers.PostDelete)

	r.Run()
}
