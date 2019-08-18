package storage

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/theonlyjohnny/phoenix/pkg/models"
)

func TestStore(t *testing.T) {
	t.Run("testStoreValidInstance", testStoreValidInstance)
	t.Run("testStoreInvalidInstance", testStoreInvalidInstance)
	t.Run("testStoreValidCluster", testStoreValidCluster)
	t.Run("testStoreInvalidCluster", testStoreInvalidCluster)
	t.Run("testStoreOther", testStoreOther)
}

func testStoreValidInstance(t *testing.T) {
	storage, err := NewLocalStorage()
	assert.NoError(t, err)

	instance := models.NewInstance("")

	err = storage.Store(InstanceEntityType, instance.PhoenixID, instance)
	assert.NoError(t, err)
}

func testStoreInvalidInstance(t *testing.T) {
	storage, err := NewLocalStorage()
	assert.NoError(t, err)

	err = storage.Store(InstanceEntityType, "", struct{}{})
	assert.Error(t, err)
}

func testStoreValidCluster(t *testing.T) {
	storage, err := NewLocalStorage()
	assert.NoError(t, err)

	cluster := &models.Cluster{}

	err = storage.Store(ClusterEntityType, cluster.Name, cluster)
	assert.NoError(t, err)
}

func testStoreInvalidCluster(t *testing.T) {
	storage, err := NewLocalStorage()
	assert.NoError(t, err)

	err = storage.Store(ClusterEntityType, "", struct{}{})
	assert.Error(t, err)
}

func testStoreOther(t *testing.T) {
	storage, err := NewLocalStorage()
	assert.NoError(t, err)

	err = storage.Store("", "", struct{}{})
	assert.Error(t, err)
}

func TestGet(t *testing.T) {
	t.Run("testGetInstance", testGetInstance)
	t.Run("testGetInvalidInstance", testGetInvalidInstance)
	t.Run("testGetCluster", testGetCluster)
	t.Run("testGetInvalidCluster", testGetInvalidCluster)
	t.Run("testGetOther", testGetOther)
}

func testGetInstance(t *testing.T) {
	storage, err := NewLocalStorage()
	assert.NoError(t, err)

	instance := models.NewInstance("")

	err = storage.Store(InstanceEntityType, instance.PhoenixID, instance)
	assert.NoError(t, err)

	res, err := storage.Get(InstanceEntityType, instance.PhoenixID)

	assert.NoError(t, err)

	returnedInstance, ok := res.(*models.Instance)

	assert.True(t, ok)
	assert.Equal(t, instance, returnedInstance)
}

func testGetInvalidInstance(t *testing.T) {
	storage, err := NewLocalStorage()
	assert.NoError(t, err)

	instance := models.NewInstance("")

	err = storage.Store(InstanceEntityType, instance.PhoenixID, instance)
	assert.NoError(t, err)

	res, err := storage.Get(InstanceEntityType, "__"+instance.PhoenixID)

	assert.Error(t, err)
	assert.Nil(t, res)
}

func testGetCluster(t *testing.T) {
	storage, err := NewLocalStorage()
	assert.NoError(t, err)

	cluster := &models.Cluster{}

	err = storage.Store(ClusterEntityType, cluster.Name, cluster)
	assert.NoError(t, err)

	res, err := storage.Get(ClusterEntityType, cluster.Name)

	assert.NoError(t, err)

	returnedCluster, ok := res.(*models.Cluster)

	assert.True(t, ok)
	assert.Equal(t, cluster, returnedCluster)
}

func testGetInvalidCluster(t *testing.T) {
	storage, err := NewLocalStorage()
	assert.NoError(t, err)

	cluster := &models.Cluster{}

	err = storage.Store(ClusterEntityType, cluster.Name, cluster)
	assert.NoError(t, err)

	res, err := storage.Get(ClusterEntityType, "__"+cluster.Name)

	assert.Error(t, err)
	assert.Nil(t, res)

}

func testGetOther(t *testing.T) {
	storage, err := NewLocalStorage()
	assert.NoError(t, err)

	cluster := &models.Cluster{}

	err = storage.Store(ClusterEntityType, cluster.Name, cluster)
	assert.NoError(t, err)
	res, err := storage.Get("", cluster.Name)

	assert.Error(t, err)
	assert.Nil(t, res)
}

func TestDelete(t *testing.T) {
	t.Run("testDeleteCluster", testDeleteCluster)
	t.Run("testDeleteInstance", testDeleteInstance)
	t.Run("testDeleteOther", testDeleteOther)
}

func testDeleteCluster(t *testing.T) {

	storage, err := NewLocalStorage()
	assert.NoError(t, err)

	cluster := &models.Cluster{}

	err = storage.Store(ClusterEntityType, cluster.Name, cluster)
	assert.NoError(t, err)

	res, err := storage.Get(ClusterEntityType, cluster.Name)

	assert.NoError(t, err)

	returnedCluster, ok := res.(*models.Cluster)

	assert.True(t, ok)
	assert.Equal(t, cluster, returnedCluster)

	err = storage.Delete(ClusterEntityType, cluster.Name)
	assert.NoError(t, err)

	res, err = storage.Get(ClusterEntityType, cluster.Name)

	assert.Error(t, err)

}

func testDeleteInstance(t *testing.T) {

	storage, err := NewLocalStorage()
	assert.NoError(t, err)

	instance := &models.Instance{}

	err = storage.Store(InstanceEntityType, instance.Name, instance)
	assert.NoError(t, err)

	res, err := storage.Get(InstanceEntityType, instance.Name)

	assert.NoError(t, err)

	returnedInstance, ok := res.(*models.Instance)

	assert.True(t, ok)
	assert.Equal(t, instance, returnedInstance)

	err = storage.Delete(InstanceEntityType, instance.Name)
	assert.NoError(t, err)

	res, err = storage.Get(InstanceEntityType, instance.Name)

	assert.Error(t, err)
}

func testDeleteOther(t *testing.T) {

	storage, err := NewLocalStorage()
	assert.NoError(t, err)

	instance := &models.Instance{}

	err = storage.Store(InstanceEntityType, instance.Name, instance)
	assert.NoError(t, err)

	res, err := storage.Get(InstanceEntityType, instance.Name)

	assert.NoError(t, err)

	returnedInstance, ok := res.(*models.Instance)

	assert.True(t, ok)
	assert.Equal(t, instance, returnedInstance)

	err = storage.Delete("", instance.Name)
	assert.Error(t, err)
}

func TestList(t *testing.T) {
	t.Run("testListInstances", testListInstances)
	t.Run("testListClusters", testListClusters)
	t.Run("testListOther", testListOther)
}

func testListInstances(t *testing.T) {

	cnt := 5
	tests := make([]*models.Instance, cnt, cnt)

	storage, err := NewLocalStorage()
	assert.NoError(t, err)

	for i := 0; i < cnt; i++ {
		tests[i] = models.NewInstance(strconv.FormatInt(int64(i), 10))

		err = storage.Store(InstanceEntityType, tests[i].PhoenixID, tests[i])
		assert.NoError(t, err)
	}

	res, err := storage.List(InstanceEntityType)
	assert.NoError(t, err)

	instances := make([]*models.Instance, len(res), len(res))
	for i, e := range res {
		instance, ok := e.(*models.Instance)
		assert.True(t, ok)
		instances[i] = instance
	}

	assert.Equal(t, tests, instances)

}

func testListClusters(t *testing.T) {

	cnt := 5
	tests := make([]*models.Cluster, cnt, cnt)

	storage, err := NewLocalStorage()
	assert.NoError(t, err)

	for i := 0; i < cnt; i++ {
		tests[i] = &models.Cluster{Name: strconv.FormatInt(int64(i), 10)}

		err = storage.Store(ClusterEntityType, tests[i].Name, tests[i])
		assert.NoError(t, err)
	}

	res, err := storage.List(ClusterEntityType)
	assert.NoError(t, err)

	clusters := make([]*models.Cluster, len(res), len(res))
	for i, e := range res {
		cluster, ok := e.(*models.Cluster)
		assert.True(t, ok)
		clusters[i] = cluster
	}

	assert.Equal(t, tests, clusters)

}

func testListOther(t *testing.T) {
	storage, err := NewLocalStorage()
	assert.NoError(t, err)

	res, err := storage.List("")

	assert.Error(t, err)
	assert.Nil(t, res)

}
