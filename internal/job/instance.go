package job

func (m *Manager) AddInstanceEvent(phoenixID string) {
	if _, err := m.cloud.GetCloudProvider(phoenixID, nil); err != nil {
		log.Errorf("Failed to load cloud provider for %s -- %s", phoenixID, err.Error())
	}
	// go m.instanceLogic.Scale(phoenixID)
}
