package deletegif

import (
	"github.com/gin-gonic/gin"
	r "gopi/internal/lib/response"
	"net/http"
)

const (
	ServerError    = http.StatusInternalServerError
	StatusNotFound = http.StatusNotFound
	StatusOK       = http.StatusOK
)

type DeleteGif interface {
	DeleteGif(id string) error
}

func New(delete DeleteGif) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		err := delete.DeleteGif(id)
		if err != nil {
			if err.Error() == "gif not found for id" {
				c.JSON(StatusNotFound, r.GifNotFound)
				return
			}
			c.JSON(ServerError, r.GifNotFound)
			return
		}

		c.JSON(StatusOK, r.OK())
	}
}
