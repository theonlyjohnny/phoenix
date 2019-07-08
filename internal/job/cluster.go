package job

func (m *Manager) addClusterEvent(clusterName string) {
	go m.operateOnClusterEvent(clusterName)
}

func (m *Manager) operateOnClusterEvent(clusterName string) {}
