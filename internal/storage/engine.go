package storage

import (
	"fmt"

	"github.com/theonlyjohnny/phoenix/internal/job"
	"github.com/theonlyjohnny/phoenix/internal/log"
	extern_storage "github.com/theonlyjohnny/phoenix/pkg/storage"
)

type Storage struct {
	backing extern_storage.Storage
	manager *job.Manager
}

func NewStorageEngine(storageType string, manager *job.Manager) (*Storage, error) {
	backing, err := getBackingByType(storageType)
	if err != nil {
		return nil, err
	}
	return &Storage{
		backing: backing,
		manager: manager,
	}, nil
}

//GetStorageByType returns an instantiated version of the requested storage
func getBackingByType(storageType string) (extern_storage.Storage, error) {

	var storage extern_storage.Storage
	var err error

	switch storageType {
	case "local":
		storage, err = extern_storage.NewLocalStorage()
	default:
		log.Errorf("Unable to find storage with type %s", storageType)
		return storage, fmt.Errorf("unknown storage %s", storage)
	}

	if err != nil {
		return storage, err
	}

	return storage, nil
}