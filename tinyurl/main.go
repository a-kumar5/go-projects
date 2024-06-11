package main

import (
	"net/http"

	_ "github.com/a-kumar5/go-projects/tinyurl/docs"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// url represents tiny url schema.
type url struct {
	Long  string `json:"long"`
	Short string `json:"short"`
}

type lurl struct {
	Long string `json:"long"`
}

type surl struct {
	Short string `json:"short"`
}

var urls = []url{
	{Long: "http://www.google.com", Short: "go.com"},
}

func main() {
	router := gin.Default()
	router.GET("/ping", getHealth)
	router.POST("/create-url", createUrl)
	router.GET("/:shorturl", getShortUrl)
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.Run("localhost:8080")
}

// PingExample godoc
// @Summary create tiny url
// @Schemes
// @Description create tiny url
// @Accept json
// @Produce json
// @Success 201
// @Router /create-url [post]
func createUrl(c *gin.Context) {
	var l lurl
	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&l); err != nil {
		return
	}
	var newUrl url
	s := l.createShortUrl()
	newUrl.Long = l.Long
	newUrl.Short = s.Short
	urls = append(urls, newUrl)
	c.IndentedJSON(http.StatusCreated, newUrl)
}

func (l *lurl) createShortUrl() surl {
	var shortUrl surl
	shortUrl = surl{Short: "xyz.com"}
	return shortUrl
}

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Accept json
// @Produce json
// @Success 200
// @Router /ping [get]
func getHealth(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "Healthy")
}

// PingExample godoc
// @Summary get tiny url
// @Schemes
// @Description Permanent Redirect
// @Accept json
// @Produce json
// @Success 301 Permanent redirect
// @Router /{shortUrl} [get]
func getShortUrl(c *gin.Context) {
	shorturl := c.Param("shorturl")
	for _, url := range urls {
		if url.Short == shorturl {
			//c.IndentedJSON(http.StatusOK, url.Long)
			c.Redirect(http.StatusMovedPermanently, url.Long)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "url not found"})
}
