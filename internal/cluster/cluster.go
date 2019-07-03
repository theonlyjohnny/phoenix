package cluster

import (
	"fmt"

	"github.com/theonlyjohnny/phoenix/internal/instance"
)

//A Cluster represents a scaling set of Instances
type Cluster struct {
	Name string `json:"name"`
}

//HasInstance returns whether or not the specified Instance is a part of this Cluster
func (c *Cluster) HasInstance(i *instance.Instance) bool {
	return i.ClusterName == c.Name
}

func (c Cluster) String() string {
	return fmt.Sprintf("Cluster[%s]{}", c.Name)
}
