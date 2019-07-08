package instance

import (
	"fmt"
	"time"
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

	UpdatedDTTM time.Time `json:"updated_dttm"`
}

func (i Instance) String() string {
	return fmt.Sprintf("Instance{PhoenixID: %s, ExternalID: %s, Name: %s, Hostname, %s, ClusterName: %s, UpdatedDTTM: %s}", i.PhoenixID, i.ExternalID, i.Name, i.Hostname, i.ClusterName, i.UpdatedDTTM)
}
