package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/theonlyjohnny/phoenix/internal/cluster"
	"github.com/theonlyjohnny/phoenix/internal/job"
	"github.com/theonlyjohnny/phoenix/internal/storage"
)

func postClusterHandler(c *gin.Context, storage *storage.Engine, manager *job.Manager) {
	var cluster cluster.Cluster
	if err := c.ShouldBindJSON(&cluster); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	storage.StoreCluster(&cluster)

	c.JSON(http.StatusOK, gin.H{})
	manager.AddClusterEvent(cluster.Name)
}
