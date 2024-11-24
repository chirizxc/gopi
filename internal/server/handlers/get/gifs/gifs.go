package gifs

import (
	"fmt"
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
		fmt.Println("GET /gifs called") // Лог вызова маршрута
		aliases, err := aliasGetter.GetAllAliases()
		if err != nil {
			fmt.Println("Error fetching aliases:", err) // Лог ошибок
			c.JSON(ServerError, r.AliasNotFound)
			return
		}
		fmt.Println("Aliases fetched:", aliases) // Лог успешного выполнения
		c.JSON(StatusOK, gin.H{"aliases": aliases})
	}

}
