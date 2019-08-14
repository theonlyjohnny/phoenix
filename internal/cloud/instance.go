package cloud

import (
	"fmt"

	"github.com/theonlyjohnny/phoenix/pkg/models"
)

func (e *Engine) GetAllInstances() (models.InstanceList, error) {
	var out models.InstanceList
	var sumErr string

	for clusterName, provider := range e.providerCache {
		list, err := provider.GetAllInstances()
		if err == nil {
			out = append(out, list...)
		} else {
			sumErr = sumErr + fmt.Sprintf(" (%s): %s", clusterName, err.Error())
		}
	}

	var err error

	if sumErr != "" {
		err = fmt.Errorf("Failed to get some instances:%s", sumErr)
	}

	return out, err
}

func (e *Engine) CreateInstance(clusterName string, newInstance *models.Instance, cmds []string) error {
	provider, err := e.GetCloudProvider(clusterName, nil)
	if err != nil {
		return err
	}
	return provider.CreateInstance(newInstance, cmds)
}
