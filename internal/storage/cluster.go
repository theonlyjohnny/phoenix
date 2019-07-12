package storage

import (
	"fmt"

	"github.com/theonlyjohnny/phoenix/internal/cluster"
	"github.com/theonlyjohnny/phoenix/pkg/storage"
)

func (s *Engine) ListClusters() (cluster.List, error) {
	vals, err := s.backing.List(storage.ClusterEntityType)
	res := make(cluster.List, len(vals))

	if err != nil {
		return res, err
	}

	for i, v := range vals {
		cluster, ok := v.(*cluster.Cluster)
		if ok {
			res[i] = cluster
		} else {
			log.Warnf("Unable to translate stored cluster to *cluster.Cluster -- %#v", v)
		}
	}
	return res, nil
}

func (s *Engine) StoreCluster(c *cluster.Cluster) error {
	err := s.backing.Store(storage.ClusterEntityType, c.Name, c)
	if err != nil {
		return err
	}
	log.Debugf("storing %s", c)

	return nil
}

func (s *Engine) GetCluster(clusterName string) (*cluster.Cluster, error) {
	v, err := s.backing.Get(storage.ClusterEntityType, clusterName)
	if err != nil {
		return nil, err
	}
	cluster, ok := v.(*cluster.Cluster)
	if !ok {
		return nil, fmt.Errorf("Unable to coerce %#v into *cluster.Cluster", v)
	}

	return cluster, nil
}
