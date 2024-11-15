package gifs

import (
	"github.com/gin-gonic/gin"
	r "gopi/internal/lib/response"
	"net/http"
)

type AliasGetter interface {
	GetAllAliases() ([]string, error)
}

func New(aliasGetter AliasGetter) gin.HandlerFunc {
	return func(c *gin.Context) {
		aliases, err := aliasGetter.GetAllAliases()
		if err != nil {
			c.JSON(http.StatusInternalServerError, r.AliasNotFound)
			return
		}

		c.JSON(http.StatusOK, gin.H{"aliases": aliases})
	}
}
