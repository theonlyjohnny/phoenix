package scale

import (
	"fmt"

	"github.com/theonlyjohnny/phoenix/internal/cloud"
	"github.com/theonlyjohnny/phoenix/internal/storage"
	"github.com/theonlyjohnny/phoenix/pkg/models"
)

var (
	installAgentCmds = []string{
		"curl -sf https://5da46805.ngrok.io/dist/phoenix-agent -o /usr/local/bin/phoenix-agent && chmod +x /usr/local/bin/phoenix-agent && phoenix-agent",
	}
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

	l.createInstances(cluster, instances)
}

func (l *ClusterLogic) createInstances(cluster *models.Cluster, instances models.InstanceList) {
	var clusterInstances models.InstanceList
	for _, i := range instances {
		if cluster.HasInstance(i) {
			clusterInstances = append(clusterInstances, i)
		}
	}
	present := len(clusterInstances)
	required := cluster.MinHealthy - present
	if required > 0 {
		log.Infof("Cluster %s Scale up -- %d < %d", cluster.Name, present, cluster.MinHealthy)
		for i := 0; i < required; i++ {
			name := fmt.Sprintf("usw1-%s-00%d", cluster.Name, i)
			if err := l.cloud.CreateInstance(cluster.Name, models.NewInstance(name), installAgentCmds); err != nil {
				log.Errorf("unable to create instance %s -- %s", name, err.Error())
			}
		}
	}
}
