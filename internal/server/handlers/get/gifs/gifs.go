package gifs

import (
	"github.com/gin-gonic/gin"
	r "gopi/internal/lib/response"
	"net/http"
)

const (
	ServerError = http.StatusInternalServerError
	StatusOK    = http.StatusOK
)

type AliasGetter interface {
	GetAllAliases() ([]string, error)
}

func New(aliasGetter AliasGetter) gin.HandlerFunc {
	return func(c *gin.Context) {
		aliases, err := aliasGetter.GetAllAliases()
		if err != nil {
			c.JSON(ServerError, r.AliasNotFound)
			return
		}

		c.JSON(StatusOK, gin.H{"aliases": aliases})
	}
}
