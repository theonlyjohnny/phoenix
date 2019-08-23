package storage

import (
	"github.com/theonlyjohnny/phoenix/pkg/models"
)

func (s *Engine) ListClusters() (models.ClusterList, error) {
	res, err := s.backing.ListClusters()

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *Engine) StoreCluster(c *models.Cluster) error {
	err := s.backing.StoreCluster(c.Name, c)
	if err != nil {
		return err
	}

	return nil
}

func (s *Engine) GetCluster(clusterName string) (*models.Cluster, error) {
	cluster, err := s.backing.GetCluster(clusterName)

	if err != nil {
		return nil, err
	}

	return cluster, nil
}
