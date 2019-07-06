package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type handler func(*gin.Context, *storageWrapper, *managerWrapper)

func unWrapHandler(realHandler handler) func(c *gin.Context) {
	return func(c *gin.Context) {
		s, ok := c.Get(StorageKey)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid Request Storage"})
		}

		sWrapper, ok := s.(storageWrapper)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid Storage"})
			return
		}

		m, ok := c.Get(ManagerKey)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid Request Manager"})
		}

		mWrapper, ok := m.(managerWrapper)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid Manager"})
			return
		}

		realHandler(c, &sWrapper, &mWrapper)
	}
}
