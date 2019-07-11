package cluster

import (
	"fmt"

	"github.com/theonlyjohnny/phoenix/internal/instance"
)

//A List is a list of Cluster pointers
type List []*Cluster

//A Cluster represents a scaling set of Instances
type Cluster struct {
	Name              string `json:"name" binding:"required"`
	MinHealthy        int    `json:"min_healthy" binding:"required"`
	CloudProviderType string `json:"cloud_provider" binding:"required"`
}

//HasInstance returns whether or not the specified Instance is a part of this Cluster
func (c *Cluster) HasInstance(i *instance.Instance) bool {
	return i.ClusterName == c.Name
}

func (c Cluster) String() string {
	return fmt.Sprintf("Cluster[%s]{MinHealthy:%d,CloudProviderType:%s}", c.Name, c.MinHealthy, c.CloudProviderType)
}
