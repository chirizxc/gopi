package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthMiddleware struct {
	User string
	Pass string
}

func New(user, pass string) gin.HandlerFunc {
	return func(c *gin.Context) {
		providedUser, providedPass, hasAuth := c.Request.BasicAuth()
		if !hasAuth || providedUser != user || providedPass != pass {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized access",
			})
			return
		}
		c.Next()
	}
}
