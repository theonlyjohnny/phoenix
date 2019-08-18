package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
	cluster := &Cluster{
		Name: testClusterName,
	}

	instance := &Instance{
		ClusterName: testClusterName,
	}

	assert.True(t, cluster.HasInstance(instance))
}

func testHasInstanceFalse(t *testing.T) {
	cluster := &Cluster{
		Name: testClusterName,
	}

	instance := &Instance{
		ClusterName: testClusterName + "_",
	}

	assert.False(t, cluster.HasInstance(instance))
}

func TestClusterString(t *testing.T) {
	cluster := Cluster{}

	assert.NotEmpty(t, cluster.String())
}
