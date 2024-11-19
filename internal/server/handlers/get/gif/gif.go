package gif

import (
	"github.com/gin-gonic/gin"
	r "gopi/internal/lib/response"
	"net/http"
	"os"
	"path/filepath"
)

const (
	StatusNotFound = http.StatusNotFound
)

type GifGetter interface {
	GetGifByAliasOrUUID(id string) (string, error)
}

func New(gifGetter GifGetter) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		path, err := gifGetter.GetGifByAliasOrUUID(id)
		if err != nil {
			c.JSON(StatusNotFound, r.GifNotFound)
			return
		}

		if _, err := os.Stat(path); os.IsNotExist(err) {
			c.JSON(StatusNotFound, r.ServerGifNotFound)
			return
		}

		c.File(filepath.Clean(path))
	}
}
