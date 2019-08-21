package job

import "fmt"

func (m *Manager) AddInstanceEvent(phoenixID string) {

	go func() {
		m.UpdateInstances()
		// m.instanceLogic.Scale(phoenixID)
	}()
}

//UpdateInstances checks the Manager's cloud provider and storage and updates the stored list of instances
func (m *Manager) UpdateInstances() error {
	s := m.storage
	c := m.cloud
	allInstances, err := c.GetAllInstances()

	if err != nil {
		return fmt.Errorf("Couldn't get all new instances -- %s", err.Error())
	}
	oldInstances, err := s.ListInstances()
	if err != nil {
		return fmt.Errorf("Couldn't get all old instances -- %s", err.Error())
	}
	delta := m.mergeInstances(allInstances, oldInstances)

	for _, oldPhoenixID := range delta.deadPhoenixIDs {
		if err := s.DeleteInstance(oldPhoenixID); err != nil {
			log.Errorf("Unable to delete instance %s from storage -- %s", oldPhoenixID, err.Error())
		}
	}

	for _, instance := range delta.instanceUpdates {
		if err := s.StoreInstance(instance); err != nil {
			log.Errorf("Unable to store instance %s to storage -- %s", instance, err.Error())
		}
	}
	return nil
}
