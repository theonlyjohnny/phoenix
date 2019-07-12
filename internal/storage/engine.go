package storage

import (
	"fmt"

	logger "github.com/theonlyjohnny/phoenix/internal/log"
	extern_storage "github.com/theonlyjohnny/phoenix/pkg/storage"
)

var log logger.Logger

func init() {
	log = logger.Log
}

type Engine struct {
	backing extern_storage.Storage
}

func NewStorageEngine(storageType string) (*Engine, error) {
	backing, err := getBackingByType(storageType)
	if err != nil {
		return nil, err
	}
	return &Engine{
		backing: backing,
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
