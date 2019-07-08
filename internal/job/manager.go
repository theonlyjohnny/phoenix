package job

import "github.com/theonlyjohnny/phoenix/internal/log"

//A Manager receives Events, recalculates state, and then applies any differences
type Manager struct{}

//NewManager returns a pointer to a newly instantiated Manager
func NewManager() (*Manager, error) {
	return &Manager{}, nil
}

//AddEvent is used to tell the Manager that there is a new Event it needs to care about
func (m *Manager) AddEvent(event Event) {
	log.Debugf("AddEvent: %#v", event)
	switch event.Type {
	case ClusterEventType:
		m.addClusterEvent(event.Key)
	default:
		log.Warnf("Unhandled event type: %s", event)
	}
}

//TODO manage # of concurrect goroutines? <- limiting
//TODO make events go into queue and cancelable via context.Context
