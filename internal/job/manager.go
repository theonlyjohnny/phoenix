package job

import "github.com/theonlyjohnny/phoenix/internal/log"

type eventType string

//An Event is fired everytime a resource changes and needs to have its state recalculated
type Event struct {
	Type eventType
	Key  string
}

//A Manager receives Events, recalculates state, and then applies any differences
type Manager struct{}

const (
	//ClusterEventType is used when a Cluster-level change has occured
	ClusterEventType = eventType("Cluster")
	//InstanceEventType is used when a Instance-level change has occured
	InstanceEventType = eventType("Instance")
)

//NewManager returns a pointer to a newly instantiated Manager
func NewManager() (*Manager, error) {
	return &Manager{}, nil
}

//AddEvent is used to tell the Manager that there is a new Event it needs to care about
func (m *Manager) AddEvent(event Event) {
	log.Debugf("AddEvent: %#v", event)
}
