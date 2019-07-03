package storage

import (
	"sync"

	"github.com/theonlyjohnny/phoenix/internal/cluster"
	"github.com/theonlyjohnny/phoenix/internal/instance"
)

var (
	instanceCache map[string]*instance.Instance
	clusterCache  map[string]*cluster.Cluster
	mutex         sync.RWMutex
)

//LocalStorage is an in-memory Storage implementation -- DO NOT USE IN PRODUCTION
type LocalStorage struct {
	instanceCache *map[string]*instance.Instance
	clusterCache  *map[string]*cluster.Cluster
	mutex         *sync.RWMutex
}

func init() {
	instanceCache = make(map[string]*instance.Instance)
	clusterCache = make(map[string]*cluster.Cluster)
}

func newLocalStorage() (*LocalStorage, error) {
	return &LocalStorage{
		&instanceCache,
		&clusterCache,
		&mutex,
	}, nil
}

func (s LocalStorage) StoreInstance(i *instance.Instance) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	pkey := i.PhoenixID
	(*s.instanceCache)[pkey] = i
	return nil
}

func (s LocalStorage) DeleteInstance(phoenixID string) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	delete(*s.instanceCache, phoenixID)
	return nil
}

func (s *LocalStorage) GetAllInstances() []*instance.Instance {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	instances := []*instance.Instance{}
	for _, instance := range *s.instanceCache {
		instances = append(instances, instance)
	}
	return instances
}

func (s LocalStorage) ListClusters() []*cluster.Cluster {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	clusters := []*cluster.Cluster{}
	for _, cluster := range *s.clusterCache {
		clusters = append(clusters, cluster)
	}
	return clusters
}

func (s LocalStorage) StoreCluster(cluster *cluster.Cluster) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	pkey := cluster.Name
	(*s.clusterCache)[pkey] = cluster
	return nil
}
