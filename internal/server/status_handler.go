package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/theonlyjohnny/phoenix/internal/job"
	"github.com/theonlyjohnny/phoenix/internal/storage"
	"github.com/theonlyjohnny/phoenix/pkg/models"
)

func postStatusHandler(c *gin.Context, storage *storage.Engine, manager *job.Manager) {
	var status models.Status
	if err := c.BindJSON(&status); err != nil {
		return
	}

	if err := c.BindHeader(&status); err != nil {
		return
	}

	if status.PhoenixID == nil {
		c.JSON(400, gin.H{"code": 400, "error": "X-Phoenix-Id header is required"})
		return
	}

	instance, err := storage.GetInstance(*status.PhoenixID)
	if err != nil {
		log.Errorf("Unable to get instance from storage on POST /status -- %s", err.Error())
		c.JSON(404, gin.H{"code": 404, "message": "Unknown phoenix_id"})
		return
	}

	instance.Status = &status
	if err := storage.StoreInstance(instance); err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{})
	manager.AddInstanceEvent(*status.PhoenixID)
}
