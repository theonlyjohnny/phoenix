package storage_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/theonlyjohnny/phoenix/internal/config"
	"github.com/theonlyjohnny/phoenix/pkg/models"
	"github.com/theonlyjohnny/phoenix/pkg/storage"
	"github.com/theonlyjohnny/phoenix/pkg/storage/redis"
	"github.com/theonlyjohnny/phoenix/testsupport"
)

type test struct {
	name string
	run  func(*testing.T, storage.Storage)
}

type storageGenFunc func(config.ComponentConfig) (storage.Storage, error)

type storageTest struct {
	name string
	//storageFunc is called to create a new struct that fulfills the Storage interface
	storageFunc storageGenFunc
	//all the functinos in setup are called in each test before calling any methods on the instantiated Storage
	cfg func() config.ComponentConfig
}

func noOpCfg() config.ComponentConfig {
	return config.ComponentConfig{}
}

var (
	storages = []storageTest{
		{
			"redis",
			redis.NewRedisStorage,
			redis.CreateTestDB,
		},
		{
			"mock",
			testsupport.NewMockStorage,
			noOpCfg,
		},
	}
	cfg = config.ComponentConfig{}
)

func TestMain(t *testing.T) {
	t.Helper()

	for _, storageSetup := range storages {
		tests := []test{
			{
				"testStoreInstance",
				testStoreInstance,
			},
			{
				"testStoreCluster",
				testStoreCluster,
			},
			{
				"testGetInstance",
				testGetInstance,
			},
			{
				"testGetCluster",
				testGetCluster,
			},
			{
				"testDeleteInstance",
				testDeleteInstance,
			},
			{
				"testDeleteCluster",
				testDeleteCluster,
			},
			{
				"testListInstances",
				testListInstances,
			},
			{
				"testListClusters",
				testListClusters,
			},
		}

		for _, test := range tests {
			testName := fmt.Sprintf("%s/%s", storageSetup.name, test.name)
			t.Run(testName, func(t *testing.T) {
				storage, err := storageSetup.storageFunc(storageSetup.cfg())

				if err != nil {
					t.Fatalf("Unable to setup %s storage -- %s", storageSetup.name, err.Error())
					return
				}
				test.run(t, storage)
			})
		}
	}
}

func testStoreInstance(t *testing.T, storage storage.Storage) {

	instance := testsupport.NewUniqueInstance()

	err := storage.StoreInstance(instance.PhoenixID, instance)
	assert.NoError(t, err)
}

func testStoreCluster(t *testing.T, storage storage.Storage) {

	cluster := &models.Cluster{}

	err := storage.StoreCluster(cluster.Name, cluster)
	assert.NoError(t, err)
}

func testGetInstance(t *testing.T, storage storage.Storage) {
	instance := models.NewInstance("")

	err := storage.StoreInstance(instance.PhoenixID, instance)
	assert.NoError(t, err)

	returnedInstance, err := storage.GetInstance(instance.PhoenixID)

	assert.NoError(t, err)
	assert.Equal(t, instance, returnedInstance)
}

func testGetCluster(t *testing.T, storage storage.Storage) {

	cluster := &models.Cluster{}

	err := storage.StoreCluster(cluster.Name, cluster)
	assert.NoError(t, err)

	returnedCluster, err := storage.GetCluster(cluster.Name)

	assert.NoError(t, err)
	assert.Equal(t, cluster, returnedCluster)
}

func testDeleteCluster(t *testing.T, storage storage.Storage) {

	cluster := &models.Cluster{}

	err := storage.StoreCluster(cluster.Name, cluster)
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

func testDeleteInstance(t *testing.T, storage storage.Storage) {
	instance := models.NewInstance("")

	err := storage.StoreInstance(instance.Name, instance)
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

func testListInstances(t *testing.T, storage storage.Storage) {
	cnt := 5
	tests := make([]*models.Instance, cnt, cnt)

	for i := 0; i < cnt; i++ {
		tests[i] = models.NewInstance(strconv.FormatInt(int64(i), 10))

		err := storage.StoreInstance(tests[i].PhoenixID, tests[i])
		assert.NoError(t, err)
	}

	instances, err := storage.ListInstances()
	assert.NoError(t, err)

	assert.ElementsMatch(t, tests, instances)

}

func testListClusters(t *testing.T, storage storage.Storage) {
	cnt := 5
	tests := make([]*models.Cluster, cnt, cnt)

	for i := 0; i < cnt; i++ {
		tests[i] = &models.Cluster{Name: strconv.FormatInt(int64(i), 10)}

		err := storage.StoreCluster(tests[i].Name, tests[i])
		assert.NoError(t, err)
	}

	clusters, err := storage.ListClusters()
	assert.NoError(t, err)

	assert.ElementsMatch(t, tests, clusters)

}
