package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/theonlyjohnny/phoenix/internal/cluster"
	"github.com/theonlyjohnny/phoenix/internal/job"
)

type postClusterRequest struct {
	ClusterName string `json:"cluster_name" binding:"required"`
}

func postClusterHandler(c *gin.Context, sw *storageWrapper, mw *managerWrapper) {
	var json postClusterRequest
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	storage := *sw.storage
	newCluster := &cluster.Cluster{Name: json.ClusterName}

	storage.StoreCluster(newCluster)

	c.JSON(http.StatusOK, gin.H{})

	manager := *mw.manager
	manager.AddEvent(job.Event{
		Type: job.ClusterEventType,
		Key:  newCluster.Name,
	})
}
