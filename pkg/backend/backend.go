package backend

import (
	"fmt"

	"github.com/theonlyjohnny/phoenix/internal/config"
	"github.com/theonlyjohnny/phoenix/internal/instance"
	"github.com/theonlyjohnny/phoenix/internal/log"
)

//Backend is a shared interface for integration with external cloud providers (EC2, GCM, etc.)
type Backend interface {
	create(config.BackendConfig) (Backend, error)
	GetAllInstances() instance.List
	// UpdateInstance(*instance.Instance) error
	// UpdateInstances(*[]instance.Instance) error
}

//GetBackendByType returns an instantiated version of the requested backend
func GetBackendByType(backendType string, cfg config.BackendConfig) (Backend, error) {
	var backend Backend
	switch backendType {
	case "ec2":
		backend = EC2{}
	default:
		log.Errorf("Unable to find backend with type %s", backendType)
		return backend, fmt.Errorf("unknown backend %s", backend)
	}
	log.Debugf("creating backend with config -- %s", cfg)
	created, err := backend.create(cfg)
	return created, err
}
