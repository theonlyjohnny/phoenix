package cloud

import "github.com/theonlyjohnny/phoenix/pkg/models"

//Provider is a shared interface for integration with external cloud providers (EC2, GCM, etc.)
type Provider interface {
	GetAllInstances() (models.InstanceList, error)
	CreateInstance(instance *models.Instance, cmds []string) error
}
