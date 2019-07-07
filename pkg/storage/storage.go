package storage

import (
	"fmt"

	"github.com/theonlyjohnny/phoenix/internal/cluster"
	"github.com/theonlyjohnny/phoenix/internal/instance"
	"github.com/theonlyjohnny/phoenix/internal/log"
)

//Storage stores Clusters and Instances to an external datastore
type Storage interface {
	ListClusters() []*cluster.Cluster
	StoreCluster(*cluster.Cluster) error

	StoreInstance(*instance.Instance) error
	DeleteInstance(string) error
	GetAllInstances() []*instance.Instance
}

//GetStorageByType returns an instantiated version of the requested storage
func GetStorageByType(storageType string) (*Storage, error) {

	var storage Storage
	var err error

	switch storageType {
	case "local":
		storage, err = newLocalStorage()
	default:
		log.Errorf("Unable to find storage with type %s", storageType)
		return &storage, fmt.Errorf("unknown storage %s", storage)
	}

	if err != nil {
		return &storage, err
	}

	return &storage, nil
}
