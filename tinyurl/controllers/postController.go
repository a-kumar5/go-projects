package controllers

import (
	"net/http"

	_ "github.com/a-kumar5/go-projects/tinyurl/docs"
	"github.com/a-kumar5/go-projects/tinyurl/initializers"
	"github.com/a-kumar5/go-projects/tinyurl/models"
	"github.com/a-kumar5/go-projects/tinyurl/pkg/encoder"
	"github.com/gin-gonic/gin"
)

type lurl struct {
	Long string
}

type surl struct {
	Short string
}

// PingExample godoc
// @Summary create tiny url
// @Schemes
// @Description create tiny url
// @Accept json
// @Produce json
// @Success 201
// @Router /create-url [post]
func CreateUrl(c *gin.Context) {
	// Call BindJSON to bind the received JSON to
	// newAlbum.
	var l lurl
	c.Bind(&l)
	url := &models.Url{Long: l.Long, Short: l.createShortUrl().Short}
	DB := initializers.ConnectToDB()
	result := DB.Create(url)
	if result.Error != nil {
		c.IndentedJSON(http.StatusBadRequest, result.Error)
	}
	c.IndentedJSON(http.StatusCreated, url)
}

func (l *lurl) createShortUrl() surl {
	var shortUrl surl
	numUrl := encoder.ToNumber(l.Long)
	shortUrl = surl{Short: encoder.ToBase62(numUrl)}
	return shortUrl
}
