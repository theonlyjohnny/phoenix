package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/theonlyjohnny/phoenix/internal/job"
	"github.com/theonlyjohnny/phoenix/internal/storage"
)

type handler func(*gin.Context, *storage.Engine, *job.Manager)

func unWrapHandler(realHandler handler) func(c *gin.Context) {
	return func(c *gin.Context) {
		w, ok := c.Get(wrapperKey)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid Request Context"})
		}

		wrapper, ok := w.(wrapper)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid Request Context"})
			return
		}

		realHandler(c, wrapper.s, wrapper.m)
	}
}
