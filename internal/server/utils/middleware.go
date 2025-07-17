package utils

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Middleware(timeout time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {

		ctx, cancel := context.WithTimeout(c.Request.Context(), timeout)
		defer cancel()

		c.Request = c.Request.WithContext(ctx)

		done := make(chan struct{})

		go func() {
			c.Next()
			close(done)
		}()

		select {
		case <-ctx.Done():
			c.AbortWithStatusJSON(http.StatusGatewayTimeout, gin.H{
				"error": "request timed out",
			})
		case <-done:
		}
	}
}
