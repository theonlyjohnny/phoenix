package storage

import (
	"github.com/theonlyjohnny/phoenix/pkg/models"
)

func (s *Engine) ListInstances() (models.InstanceList, error) {
	res, err := s.backing.ListInstances()

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *Engine) DeleteInstance(key string) error {
	return s.backing.DeleteInstance(key)
}

func (s *Engine) StoreInstance(i *models.Instance) error {
	return s.backing.StoreInstance(i.PhoenixID, i)
}

func (s *Engine) GetInstance(phoenixID string) (*models.Instance, error) {
	instance, err := s.backing.GetInstance(phoenixID)
	if err != nil {
		return nil, err
	}

	return instance, nil

}
