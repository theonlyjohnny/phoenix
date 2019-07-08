package job

import "fmt"

const (
	//ClusterEventType is used when a Cluster-level change has occured
	ClusterEventType = eventType("Cluster")
	//InstanceEventType is used when a Instance-level change has occured
	InstanceEventType = eventType("Instance")
)

type eventType string

//An Event is fired everytime a resource changes and needs to have its state recalculated
type Event struct {
	Type eventType
	Key  string
}

func (e Event) String() string {
	return fmt.Sprintf("Event{Type: %s, Key: %s}", e.Type, e.Key)
}
