package cloud

import (
	"github.com/theonlyjohnny/phoenix/internal/instance"
)

//Provider is a shared interface for integration with external cloud providers (EC2, GCM, etc.)
type Provider interface {
	GetAllInstances() (instance.List, error)
	CreateInstance(instance *instance.Instance, cmds []string) error
}
