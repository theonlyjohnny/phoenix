package testsupport

import (
	"fmt"

	"github.com/theonlyjohnny/phoenix/internal/config"
	"github.com/theonlyjohnny/phoenix/pkg/models"
	"github.com/theonlyjohnny/phoenix/pkg/storage"
)

//MockStorage is an in-memory Storage implementation -- DO NOT USE IN PRODUCTION
type MockStorage struct {
	instanceCache map[string]*models.Instance
	clusterCache  map[string]*models.Cluster
}

func NewMockStorage(_ config.ComponentConfig) (storage.Storage, error) {
	instanceCache := make(map[string]*models.Instance)
	clusterCache := make(map[string]*models.Cluster)

	return &MockStorage{
		instanceCache,
		clusterCache,
	}, nil
}

func (s *MockStorage) StoreCluster(key string, v *models.Cluster) error {

	s.clusterCache[key] = v
	return nil
}

func (s *MockStorage) StoreInstance(key string, v *models.Instance) error {

	s.instanceCache[key] = v
	return nil
}

func (s *MockStorage) DeleteCluster(key string) error {
	delete(s.clusterCache, key)
	return nil
}

func (s *MockStorage) DeleteInstance(key string) error {
	delete(s.instanceCache, key)
	return nil
}

func (s *MockStorage) GetCluster(key string) (*models.Cluster, error) {

	res, ok := s.clusterCache[key]

	if !ok {
		return nil, fmt.Errorf("Invalid key %s", key)
	}

	return res, nil
}

func (s *MockStorage) GetInstance(key string) (*models.Instance, error) {

	res, ok := s.instanceCache[key]

	if !ok {
		return nil, fmt.Errorf("Invalid key %s", key)
	}

	return res, nil
}

func (s *MockStorage) ListClusters() (models.ClusterList, error) {
	res := models.ClusterList{}

	for _, cluster := range s.clusterCache {
		res = append(res, cluster)
	}

	return res, nil
}

func (s *MockStorage) ListInstances() (models.InstanceList, error) {
	res := models.InstanceList{}

	for _, instance := range s.instanceCache {
		res = append(res, instance)
	}

	return res, nil
}
