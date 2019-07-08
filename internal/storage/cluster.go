package storage

import (
	"github.com/theonlyjohnny/phoenix/internal/cluster"
	"github.com/theonlyjohnny/phoenix/internal/job"
	"github.com/theonlyjohnny/phoenix/internal/log"
	"github.com/theonlyjohnny/phoenix/pkg/storage"
)

func (s *Storage) ListClusters() (cluster.List, error) {
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

func (s *Storage) StoreCluster(c *cluster.Cluster) error {
	err := s.backing.Store(storage.ClusterEntityType, c.Name, c)
	if err != nil {
		return err
	}

	s.manager.AddEvent(job.Event{
		Type: job.ClusterEventType,
		Key:  c.Name,
	})
	return nil
}
