package instance

import (
	"fmt"
	"time"

	uuid "github.com/satori/go.uuid"
)

//A List is a list of Instance pointers
type List []*Instance

//An Instance represents a managed server provided by a Backend
type Instance struct {
	PhoenixID  string `json:"phoenix_id"`
	ExternalID string `json:"external_id"`

	Name     string `json:"name"`
	Hostname string `json:"hostname"`

	ClusterName string `json:"cluster_name"`

	Location Location `json:"location"`

	Status *Status `json:"status"`

	UpdatedDTTM time.Time `json:"updated_dttm"`
}

func (i Instance) String() string {
	return fmt.Sprintf("Instance{PhoenixID: %s, ExternalID: %s, Name: %s, Hostname, %s, ClusterName: %s, UpdatedDTTM: %s, Location: %s}", i.PhoenixID, i.ExternalID, i.Name, i.Hostname, i.ClusterName, i.UpdatedDTTM, i.Location)
}

func NewInstance(name string) *Instance {
	phoenixID := uuid.NewV4().String()
	return &Instance{
		PhoenixID: phoenixID,
		Name:      name,
	}
}
