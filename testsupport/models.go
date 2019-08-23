package testsupport

import (
	"fmt"
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/theonlyjohnny/phoenix/pkg/models"
)

func NewUniqueCluster() *models.Cluster {
	name := fmt.Sprintf("cluster_%s", uuid.NewV4().String())
	minHealthy := r.Int()
	cloudProviderType := uuid.NewV4().String()

	return &models.Cluster{
		Name:              name,
		MinHealthy:        minHealthy,
		CloudProviderType: cloudProviderType,
	}
}

func NewUniqueInstance() *models.Instance {
	phoenixID := fmt.Sprintf("phoenix_id_%s", uuid.NewV4().String())
	externalID := fmt.Sprintf("external_id_%s", uuid.NewV4().String())
	name := fmt.Sprintf("name_%s", uuid.NewV4().String())
	hostname := fmt.Sprintf("hostname_%s", uuid.NewV4().String())
	clusterName := fmt.Sprintf("clustername_%s", uuid.NewV4().String())

	region := fmt.Sprintf("region_%s", uuid.NewV4().String())
	zone := fmt.Sprintf("zone_%s", uuid.NewV4().String())

	cpuUsage := r.Float64()
	memUsage := r.Float64()
	healthy := r.Int() > 50

	updatedDTTM := time.Now()

	return &models.Instance{
		PhoenixID:   phoenixID,
		ExternalID:  externalID,
		Name:        name,
		Hostname:    hostname,
		ClusterName: clusterName,
		Location: models.Location{
			Region: region,
			Zone:   zone,
		},
		Status: &models.Status{
			CPUUsage: cpuUsage,
			MemUsage: memUsage,
			Healthy:  healthy,
		},
		UpdatedDTTM: updatedDTTM,
	}
}
