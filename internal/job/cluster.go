package job

import "github.com/theonlyjohnny/phoenix/internal/log"

func (m *Manager) AddClusterEvent(clusterName string) {
	//TODO per-cluster provider overrides
	if _, err := m.cloud.GetCloudProvider(clusterName, nil); err != nil {
		log.Errorf("Failed to load cloud provider for %s -- %s", clusterName, err.Error())
	}
	go m.clusterLogic.Scale(clusterName)
}
