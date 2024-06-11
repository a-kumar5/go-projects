package controllers

import (
	"fmt"
	"net/http"

	"github.com/a-kumar5/go-projects/tinyurl/initializers"
	"github.com/a-kumar5/go-projects/tinyurl/models"
	"github.com/gin-gonic/gin"
)

// PingExample godoc
// @Summary get tiny url
// @Schemes
// @Description Permanent Redirect
// @Accept json
// @Produce json
// @Success 301 Permanent redirect
// @Router /{shortUrl} [get]
func GetShortUrl(c *gin.Context) {
	shorturl := c.Param("shorturl")
	urlModel := &models.Url{}
	if shorturl == "" {
		c.IndentedJSON(http.StatusInternalServerError, "id cannot be empty")
		return
	}
	fmt.Println("the shorturl is ", shorturl)
	db := initializers.ConnectToDB()
	result := db.Select("long").Where("short = ?", shorturl).First(urlModel)
	if result.Error != nil {
		c.IndentedJSON(http.StatusBadRequest, result.Error)
	}
	c.Redirect(http.StatusMovedPermanently, urlModel.Long)
	return
}

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Accept json
// @Produce json
// @Success 200
// @Router /ping [get]
func GetHealth(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "Healthy")
}
