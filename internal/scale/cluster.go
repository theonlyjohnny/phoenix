package scale

import (
	"fmt"

	"github.com/theonlyjohnny/phoenix/internal/cloud"
	"github.com/theonlyjohnny/phoenix/internal/instance"
	"github.com/theonlyjohnny/phoenix/internal/storage"
)

type ClusterLogic struct {
	storage *storage.Engine
	cloud   *cloud.Engine
}

func NewClusterLogic(s *storage.Engine, c *cloud.Engine) *ClusterLogic {
	return &ClusterLogic{
		s,
		c,
	}
}

func (l *ClusterLogic) Scale(clusterName string) {
	cluster, err := l.storage.GetCluster(clusterName)
	if err != nil {
		log.Errorf("Couldn't scale %s -- %s", clusterName, err.Error())
	}
	log.Debugf("scaling %s", cluster)
	if cluster == nil {
		log.Warnf("Told to Scale a non-existent cluster? %s", clusterName)
		return
	}
	instances, err := l.storage.ListInstances()
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
	required := cluster.MinHealthy - present
	if required > 0 {
		log.Infof("Cluster %s Scale up -- %d < %d", clusterName, present, cluster.MinHealthy)
		for i := 0; i < required; i++ {
			name := fmt.Sprintf("usw1-%s-00%d", clusterName, i)
			if err := l.cloud.CreateInstance(clusterName, instance.NewInstance(name)); err != nil {
				log.Errorf("unable to create instance %s -- %s", name, err.Error())
			}
		}
	}
}
