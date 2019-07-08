package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/theonlyjohnny/phoenix/internal/storage"
)

type handler func(*gin.Context, *storage.Engine)

func unWrapHandler(realHandler handler) func(c *gin.Context) {
	return func(c *gin.Context) {
		s, ok := c.Get(StorageKey)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid Request Engine"})
		}

		storage, ok := s.(*storage.Engine)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid Engine"})
			return
		}

		realHandler(c, storage)
	}
}
