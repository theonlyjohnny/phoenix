package storage

import (
	"fmt"
	"sync"

	"github.com/theonlyjohnny/phoenix/pkg/models"
)

//LocalStorage is an in-memory Storage implementation -- DO NOT USE IN PRODUCTION
type LocalStorage struct {
	instanceCache map[string]*models.Instance
	clusterCache  map[string]*models.Cluster
	mutex         *sync.RWMutex
}

func NewLocalStorage() (Storage, error) {
	instanceCache := make(map[string]*models.Instance)
	clusterCache := make(map[string]*models.Cluster)
	var mutex sync.RWMutex

	return &LocalStorage{
		instanceCache,
		clusterCache,
		&mutex,
	}, nil
}

func (s *LocalStorage) StoreCluster(key string, v *models.Cluster) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.clusterCache[key] = v
	return nil
}

func (s *LocalStorage) StoreInstance(key string, v *models.Instance) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.instanceCache[key] = v
	return nil
}

func (s *LocalStorage) DeleteCluster(key string) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	delete(s.clusterCache, key)
	return nil
}

func (s *LocalStorage) DeleteInstance(key string) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	delete(s.instanceCache, key)
	return nil
}

func (s *LocalStorage) GetCluster(key string) (*models.Cluster, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	res, ok := s.clusterCache[key]

	if !ok {
		return nil, fmt.Errorf("Invalid key %s", key)
	}

	return res, nil
}

func (s *LocalStorage) GetInstance(key string) (*models.Instance, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	res, ok := s.instanceCache[key]

	if !ok {
		return nil, fmt.Errorf("Invalid key %s", key)
	}

	return res, nil
}

func (s *LocalStorage) ListClusters() (models.ClusterList, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	res := models.ClusterList{}

	for _, cluster := range s.clusterCache {
		res = append(res, cluster)
	}

	return res, nil
}

func (s *LocalStorage) ListInstances() (models.InstanceList, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	res := models.InstanceList{}

	for _, instance := range s.instanceCache {
		res = append(res, instance)
	}

	return res, nil
}

func (s *LocalStorage) Flush() error {
	s.instanceCache = make(map[string]*models.Instance)
	s.clusterCache = make(map[string]*models.Cluster)
	return nil
}
