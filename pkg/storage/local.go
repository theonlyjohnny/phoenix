package storage

import (
	"github.com/theonlyjohnny/phoenix/internal/cluster"
	"github.com/theonlyjohnny/phoenix/internal/instance"
)

var (
	instanceCache map[string]*instance.Instance
	clusterCache  map[string]*cluster.Cluster
)

//LocalStorage is an in-memory Storage implementation -- DO NOT USE IN PRODUCTION
type LocalStorage struct {
	instanceCache *map[string]*instance.Instance
	clusterCache  *map[string]*cluster.Cluster
}

func init() {
	instanceCache = make(map[string]*instance.Instance)
	clusterCache = make(map[string]*cluster.Cluster)
}

func newLocalStorage() (*LocalStorage, error) {
	return &LocalStorage{
		&instanceCache,
		&clusterCache,
	}, nil
}

func (s LocalStorage) StoreInstance(i *instance.Instance) error {
	pkey := i.PhoenixID
	(*s.instanceCache)[pkey] = i
	return nil
}

func (s LocalStorage) DeleteInstance(phoenixID string) error {
	delete(*s.instanceCache, phoenixID)
	return nil
}

func (s *LocalStorage) GetAllInstances() []*instance.Instance {
	instances := []*instance.Instance{}
	for _, instance := range *s.instanceCache {
		instances = append(instances, instance)
	}
	return instances
}

func (s LocalStorage) ListClusters() []*cluster.Cluster {
	clusters := []*cluster.Cluster{}
	for _, cluster := range *s.clusterCache {
		clusters = append(clusters, cluster)
	}
	return clusters
}

func (s LocalStorage) StoreCluster(cluster *cluster.Cluster) error {
	pkey := cluster.Name
	(*s.clusterCache)[pkey] = cluster
	return nil
}
