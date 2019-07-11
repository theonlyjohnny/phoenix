package cloud

import (
	"fmt"

	"github.com/theonlyjohnny/phoenix/internal/instance"
)

func (e *Engine) GetAllInstances() (instance.List, error) {
	var out instance.List
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

func (e *Engine) CreateInstance(clusterName string, newInstance *instance.Instance) error {
	provider, err := e.GetCloudProvider(clusterName, nil)
	if err != nil {
		return err
	}
	return provider.CreateInstance(newInstance)
}
