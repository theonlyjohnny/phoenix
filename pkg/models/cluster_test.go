package models_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/theonlyjohnny/phoenix/pkg/models"
)

const (
	testClusterName              = "test_cluster_name"
	testClusterMinHealthy        = 1
	testClusterCloudProviderType = "test_cluster_cloud_provider_type"
)

func TestHasInstance(t *testing.T) {
	t.Run("testHasInstanceTrue", testHasInstanceTrue)
	t.Run("testHasInstanceFalse", testHasInstanceFalse)
}

func testHasInstanceTrue(t *testing.T) {
	cluster := &models.Cluster{
		Name: testClusterName,
	}

	instance := &models.Instance{
		ClusterName: testClusterName,
	}

	assert.True(t, cluster.HasInstance(instance))
}

func testHasInstanceFalse(t *testing.T) {
	cluster := &models.Cluster{
		Name: testClusterName,
	}

	instance := &models.Instance{
		ClusterName: testClusterName + "_",
	}

	assert.False(t, cluster.HasInstance(instance))
}

func TestClusterString(t *testing.T) {
	cluster := models.Cluster{}

	assert.NotEmpty(t, cluster.String())
}
