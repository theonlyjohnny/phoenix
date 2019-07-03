package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/theonlyjohnny/phoenix/internal/cluster"
)

type postClusterRequest struct {
	ClusterName string `json:"cluster_name" binding:"required"`
}

func postClusterHandler(c *gin.Context) {
	var json postClusterRequest
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	s, ok := c.Get(StorageKey)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid Request Storage"})
		return
	}

	wrapper, ok := s.(storageWrapper)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid Storage"})
		return
	}

	fmt.Printf("wrapper: %#v \n", wrapper)
	storage := *wrapper.storage
	newCluster := &cluster.Cluster{Name: json.ClusterName}

	storage.StoreCluster(newCluster)

	c.JSON(http.StatusOK, gin.H{})
}
