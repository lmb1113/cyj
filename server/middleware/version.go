package middleware

import (
	"github.com/gin-gonic/gin"
)

func Version() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("x-version", c.Request.Header.Get("x-version"))
		c.Next()
	}
}
