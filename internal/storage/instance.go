package storage

import (
	"fmt"

	"github.com/theonlyjohnny/phoenix/pkg/models"
	"github.com/theonlyjohnny/phoenix/pkg/storage"
)

func (s *Engine) ListInstances() (models.InstanceList, error) {
	vals, err := s.backing.List(storage.InstanceEntityType)
	res := make(models.InstanceList, len(vals))
	if err != nil {
		return res, err
	}
	for i, v := range vals {
		instance, ok := v.(*models.Instance)
		if ok {
			res[i] = instance
		} else {
			log.Warnf("Unable to translate stored instance to *instance.Instance -- %#v", v)
		}
	}
	return res, nil
}

func (s *Engine) DeleteInstance(key string) error {
	return s.backing.Delete(storage.InstanceEntityType, key)
}

func (s *Engine) StoreInstance(i *models.Instance) error {
	return s.backing.Store(storage.InstanceEntityType, i.PhoenixID, i)
}

func (s *Engine) GetInstance(phoenixID string) (*models.Instance, error) {
	instanceInterface, err := s.backing.Get(storage.InstanceEntityType, phoenixID)
	if err != nil {
		return nil, err
	}

	instance, ok := instanceInterface.(*models.Instance)
	if !ok {
		return nil, fmt.Errorf("Unable to translate stored instance to *instance.Instance -- %#v", instanceInterface)
	}
	return instance, nil

}
