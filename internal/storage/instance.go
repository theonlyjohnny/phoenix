package storage

import (
	"github.com/theonlyjohnny/phoenix/internal/instance"
	"github.com/theonlyjohnny/phoenix/internal/log"
	"github.com/theonlyjohnny/phoenix/pkg/storage"
)

func (s *Engine) ListInstances() (instance.List, error) {
	vals, err := s.backing.List(storage.InstanceEntityType)
	res := make(instance.List, len(vals))
	if err != nil {
		return res, err
	}
	for i, v := range vals {
		instance, ok := v.(*instance.Instance)
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

func (s *Engine) StoreInstance(i *instance.Instance) error {
	return s.backing.Store(storage.InstanceEntityType, i.PhoenixID, i)
}
