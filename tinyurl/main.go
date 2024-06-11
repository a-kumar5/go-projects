package main

import (
	"github.com/a-kumar5/go-projects/tinyurl/controllers"
	_ "github.com/a-kumar5/go-projects/tinyurl/docs"
	"github.com/a-kumar5/go-projects/tinyurl/initializers"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	router := gin.Default()
	router.GET("/", controllers.GetHealth)
	router.POST("/create-url", controllers.CreateUrl)
	router.GET("/:shorturl", controllers.GetShortUrl)
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.Run()
}
