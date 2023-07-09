package main

import (
	"github.com/Tanmoy037/go-crud/initializers"
	"github.com/Tanmoy037/go-crud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
