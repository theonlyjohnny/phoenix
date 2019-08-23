package storage

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/theonlyjohnny/phoenix/pkg/models"
	"github.com/theonlyjohnny/phoenix/testsupport"
)

type storageGenFunc func() (Storage, error)

var storageFuncs = map[string]storageGenFunc{
	"redis": NewRedisStorage,
	"local": NewLocalStorage,
}

func TestStore(t *testing.T) {
	t.Helper()
	//TODO overwrite tests
	for name, v := range storageFuncs {
		t.Run(fmt.Sprintf("testStoreInstance/%s", name), func(t *testing.T) { testStoreInstance(t, v) })
		t.Run(fmt.Sprintf("testStoreCluster/%s", name), func(t *testing.T) { testStoreCluster(t, v) })
	}
}

func testStoreInstance(t *testing.T, storageFunc storageGenFunc) {
	storage, err := storageFunc()
	assert.NoError(t, err)

	instance := models.NewInstance("")

	err = storage.StoreInstance(instance.PhoenixID, instance)
	assert.NoError(t, err)
}

func testStoreCluster(t *testing.T, storageFunc storageGenFunc) {
	storage, err := storageFunc()
	assert.NoError(t, err)

	cluster := &models.Cluster{}

	err = storage.StoreCluster(cluster.Name, cluster)
	assert.NoError(t, err)
}

func TestGet(t *testing.T) {
	t.Helper()
	for name, v := range storageFuncs {
		t.Run(fmt.Sprintf("testGetInstance/%s", name), func(t *testing.T) { testGetInstance(t, v) })
		t.Run(fmt.Sprintf("testGetCluster/%s", name), func(t *testing.T) { testGetCluster(t, v) })

		storage, _ := v()
		storage.Flush()
	}
}

func testGetInstance(t *testing.T, storageFunc storageGenFunc) {
	storage, err := storageFunc()
	assert.NoError(t, err)

	instance := models.NewInstance("")

	err = storage.StoreInstance(instance.PhoenixID, instance)
	assert.NoError(t, err)

	returnedInstance, err := storage.GetInstance(instance.PhoenixID)

	assert.NoError(t, err)
	assert.Equal(t, instance, returnedInstance)
}

func testGetCluster(t *testing.T, storageFunc storageGenFunc) {
	storage, err := storageFunc()
	assert.NoError(t, err)

	cluster := &models.Cluster{}

	err = storage.StoreCluster(cluster.Name, cluster)
	assert.NoError(t, err)

	returnedCluster, err := storage.GetCluster(cluster.Name)

	assert.NoError(t, err)
	assert.Equal(t, cluster, returnedCluster)
}

func TestDelete(t *testing.T) {
	t.Helper()
	for name, v := range storageFuncs {
		t.Run(fmt.Sprintf("testDeleteCluster/%s", name), func(t *testing.T) { testDeleteCluster(t, v) })
		t.Run(fmt.Sprintf("testDeleteInstance/%s", name), func(t *testing.T) { testDeleteInstance(t, v) })

		storage, _ := v()
		storage.Flush()
	}
}

func testDeleteCluster(t *testing.T, storageFunc storageGenFunc) {
	storage, err := storageFunc()
	assert.NoError(t, err)

	cluster := &models.Cluster{}

	err = storage.StoreCluster(cluster.Name, cluster)
	assert.NoError(t, err)

	returnedCluster, err := storage.GetCluster(cluster.Name)

	assert.NoError(t, err)

	assert.Equal(t, cluster, returnedCluster)

	err = storage.DeleteCluster(cluster.Name)
	assert.NoError(t, err)

	postDeleteGetCluster, err := storage.GetCluster(cluster.Name)

	assert.Error(t, err)
	assert.Nil(t, postDeleteGetCluster)

}

func testDeleteInstance(t *testing.T, storageFunc storageGenFunc) {
	storage, err := storageFunc()
	assert.NoError(t, err)

	instance := models.NewInstance("")

	err = storage.StoreInstance(instance.Name, instance)
	assert.NoError(t, err)

	returnedInstance, err := storage.GetInstance(instance.Name)

	assert.NoError(t, err)

	assert.Equal(t, instance, returnedInstance)

	err = storage.DeleteInstance(instance.Name)
	assert.NoError(t, err)

	postDeleteGetInstance, err := storage.GetInstance(instance.Name)

	assert.Error(t, err)
	assert.Nil(t, postDeleteGetInstance)

}

func TestList(t *testing.T) {
	t.Helper()
	for name, v := range storageFuncs {
		t.Run(fmt.Sprintf("testListInstances/%s", name), func(t *testing.T) { testListInstances(t, v) })
		t.Run(fmt.Sprintf("testListClusters/%s", name), func(t *testing.T) { testListClusters(t, v) })

		storage, _ := v()
		storage.Flush()
	}
}

func testListInstances(t *testing.T, storageFunc storageGenFunc) {

	storage, err := storageFunc()
	assert.NoError(t, err)

	cnt := 5
	tests := make([]*models.Instance, cnt, cnt)

	for i := 0; i < cnt; i++ {
		tests[i] = models.NewInstance(strconv.FormatInt(int64(i), 10))

		err = storage.StoreInstance(tests[i].PhoenixID, tests[i])
		assert.NoError(t, err)
	}

	instances, err := storage.ListInstances()
	assert.NoError(t, err)

	assert.ElementsMatch(t, tests, instances)

}

func testListClusters(t *testing.T, storageFunc storageGenFunc) {

	storage, err := storageFunc()
	assert.NoError(t, err)

	cnt := 5
	tests := make([]*models.Cluster, cnt, cnt)

	for i := 0; i < cnt; i++ {
		tests[i] = &models.Cluster{Name: strconv.FormatInt(int64(i), 10)}

		err = storage.StoreCluster(tests[i].Name, tests[i])
		assert.NoError(t, err)
	}

	clusters, err := storage.ListClusters()
	assert.NoError(t, err)

	assert.ElementsMatch(t, tests, clusters)

}

func TestFlush(t *testing.T) {
	t.Helper()
	for name, storageFunc := range storageFuncs {
		t.Run(fmt.Sprintf("TestFlush/%s", name), func(t *testing.T) {
			storage, err := storageFunc()

			assert.NoError(t, err)

			cluster := testsupport.NewUniqueCluster()
			instance := testsupport.NewUniqueInstance()

			storage.StoreCluster(cluster.Name, cluster)
			storage.StoreInstance(instance.PhoenixID, instance)

			err = storage.Flush()
			assert.NoError(t, err)

			resCluster, err := storage.GetCluster(cluster.Name)
			assert.Error(t, err)
			assert.Nil(t, resCluster)

			resInstance, err := storage.GetInstance(instance.PhoenixID)
			assert.Error(t, err)
			assert.Nil(t, resInstance)

		})
	}
}
