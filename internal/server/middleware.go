package server

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/theonlyjohnny/phoenix/internal/storage"
)

func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}

func logMiddleware() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {

		// your custom format
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	})
}

func storageMiddleware(s *storage.Storage) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(StorageKey, s)
		c.Next()
	}
}
