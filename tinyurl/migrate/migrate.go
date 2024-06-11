package main

import (
	"github.com/a-kumar5/go-projects/tinyurl/initializers"
	"github.com/a-kumar5/go-projects/tinyurl/models"
)

func main() {
	initializers.LoadEnvVariables()
	DB := initializers.ConnectToDB()
	DB.AutoMigrate(&models.Url{})
}
