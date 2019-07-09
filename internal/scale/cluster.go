package scale

import (
	"github.com/theonlyjohnny/phoenix/internal/instance"
	"github.com/theonlyjohnny/phoenix/internal/log"
	"github.com/theonlyjohnny/phoenix/internal/storage"
)

type ClusterLogic struct {
	storage *storage.Engine
}

func NewClusterLogic(s *storage.Engine) *ClusterLogic {
	return &ClusterLogic{
		s,
	}
}

func (c *ClusterLogic) Scale(clusterName string) {
	cluster, err := c.storage.GetCluster(clusterName)
	if err != nil {
		log.Errorf("Couldn't scale %s -- %s", clusterName, err.Error())
	}
	log.Debugf("scaling %s", cluster)
	if cluster == nil {
		log.Warnf("Told to Scale a non-existent cluster? %s", clusterName)
		return
	}
	instances, err := c.storage.ListInstances()
	if err != nil {
		log.Errorf("Couldn't scale %s -- %s", cluster.Name, err.Error())
	}
	var clusterInstances instance.List
	for _, i := range instances {
		if cluster.HasInstance(i) {
			clusterInstances = append(clusterInstances, i)
		}
	}
	present := len(clusterInstances)
	if present < cluster.MinHealthy {
		log.Infof("Cluster %s Scale up -- %d < %d", clusterName, present, cluster.MinHealthy)
	}
	//todo real logic
}
