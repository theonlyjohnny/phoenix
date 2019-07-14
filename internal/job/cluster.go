package job

func (m *Manager) AddClusterEvent(clusterName string) {
	if _, err := m.cloud.GetCloudProvider(clusterName, nil); err != nil {
		log.Errorf("Failed to load cloud provider for %s -- %s", clusterName, err.Error())
	}
	go m.clusterLogic.Scale(clusterName)
}
