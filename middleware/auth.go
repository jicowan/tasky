package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/jeffthorne/tasky/auth"
)

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !auth.ValidateSession(c) {
			c.Abort()
			return
		}
		c.Next()
	}
}
