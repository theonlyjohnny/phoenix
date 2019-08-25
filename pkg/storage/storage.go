package storage

import "github.com/theonlyjohnny/phoenix/pkg/models"

//Storage stores interfaces to some (ideally remote) persistent storage
type Storage interface {
	StoreCluster(string, *models.Cluster) error //TODO remove the key from this interface as it can be inferred from the type? but then how does Get know what key to use...
	StoreInstance(string, *models.Instance) error

	ListClusters() (models.ClusterList, error)
	ListInstances() (models.InstanceList, error)

	GetCluster(string) (*models.Cluster, error)
	GetInstance(string) (*models.Instance, error)

	DeleteCluster(string) error
	DeleteInstance(string) error
}
