package storage

import (
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
