package job

func (m *Manager) AddClusterEvent(clusterName string) {
	go m.clusterLogic.Scale(clusterName)
}
