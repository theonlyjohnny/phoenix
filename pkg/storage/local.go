package storage

import (
	"fmt"
	"sync"

	"github.com/theonlyjohnny/phoenix/pkg/models"
)

var (
	instanceCache map[string]*models.Instance
	clusterCache  map[string]*models.Cluster
	mutex         sync.RWMutex
)

//LocalStorage is an in-memory Storage implementation -- DO NOT USE IN PRODUCTION
type LocalStorage struct {
	instanceCache *map[string]*models.Instance
	clusterCache  *map[string]*models.Cluster
	mutex         *sync.RWMutex
}

func init() {
	instanceCache = make(map[string]*models.Instance)
	clusterCache = make(map[string]*models.Cluster)
}

func NewLocalStorage() (LocalStorage, error) {
	return LocalStorage{
		&instanceCache,
		&clusterCache,
		&mutex,
	}, nil
}

func (s LocalStorage) Store(t EntityType, key string, v interface{}) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	switch t {
	case InstanceEntityType:
		i, ok := v.(*models.Instance)
		if !ok {
			return fmt.Errorf("Told to save instance but value was not instance")
		}
		(*s.instanceCache)[key] = i
	case ClusterEntityType:
		c, ok := v.(*models.Cluster)
		if !ok {
			return fmt.Errorf("Told to save cluster but value was not cluster")
		}
		(*s.clusterCache)[key] = c
	default:
		return fmt.Errorf("Unknown EntityType: %s", t)
	}
	return nil
}

func (s LocalStorage) Delete(t EntityType, key string) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	switch t {
	case InstanceEntityType:
		delete(*s.instanceCache, key)
	case ClusterEntityType:
		delete(*s.clusterCache, key)
	default:
		return fmt.Errorf("Unkown EntityType: %s", t)
	}
	return nil
}

func (s LocalStorage) Get(t EntityType, key string) (interface{}, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	var res interface{}
	switch t {
	case InstanceEntityType:
		res = (*s.instanceCache)[key]
	case ClusterEntityType:
		res = (*s.clusterCache)[key]
	default:
		return res, fmt.Errorf("Unkown EntityType: %s", t)
	}
	return res, nil
}

func (s LocalStorage) List(t EntityType) ([]interface{}, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	res := []interface{}{}
	switch t {
	case InstanceEntityType:
		for _, instance := range *s.instanceCache {
			res = append(res, instance)
		}
	case ClusterEntityType:
		for _, cluster := range *s.clusterCache {
			res = append(res, cluster)
		}
	default:
		return res, fmt.Errorf("Unkown EntityType: %s", t)
	}
	return res, nil
}
