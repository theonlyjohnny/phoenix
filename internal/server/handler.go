package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/theonlyjohnny/phoenix/internal/storage"
)

type handler func(*gin.Context, *storage.Storage)

func unWrapHandler(realHandler handler) func(c *gin.Context) {
	return func(c *gin.Context) {
		s, ok := c.Get(StorageKey)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid Request Storage"})
		}

		storage, ok := s.(*storage.Storage)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid Storage"})
			return
		}

		realHandler(c, storage)
	}
}
