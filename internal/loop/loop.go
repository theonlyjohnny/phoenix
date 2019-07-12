package loop

import (
	"time"

	"github.com/theonlyjohnny/phoenix/internal/cloud"
	"github.com/theonlyjohnny/phoenix/internal/config"
	"github.com/theonlyjohnny/phoenix/internal/job"
	logger "github.com/theonlyjohnny/phoenix/internal/log"
	"github.com/theonlyjohnny/phoenix/internal/storage"
)

var log logger.Logger

func init() {
	log = logger.Log
}

type phoenixLoop struct {
	loopInterval time.Duration

	cloud   *cloud.Engine
	storage *storage.Engine
	manager *job.Manager
}

// Start starts the main Phoenix loop
func Start(cfg *config.Config, s *storage.Engine, c *cloud.Engine, m *job.Manager) error {

	loop, err := newPhoenixLoop(cfg, s, c, m)
	if err != nil {
		return err
	}
	loop.start()
	return nil
}

func newPhoenixLoop(cfg *config.Config, s *storage.Engine, c *cloud.Engine, m *job.Manager) (*phoenixLoop, error) {
	return &phoenixLoop{
		cfg.LoopInterval,
		c,
		s,
		m,
	}, nil
}

func (l *phoenixLoop) start() {
	loopInterval := l.loopInterval
	log.Debugf("starting loop -- instant + every %s \n", loopInterval)
	ticker := time.NewTicker(loopInterval)

	l.tick()

	for range ticker.C {
		l.tick()
	}
}

func (l *phoenixLoop) tick() {
	log.Debugf("Tick")
	l.updateInstances()
	l.checkInstances()
}

func (l *phoenixLoop) updateInstances() {
	s := l.storage
	c := l.cloud
	allInstances, err := c.GetAllInstances()
	if err != nil {
		log.Errorf("Couldn't get all new instances -- %s", err.Error())
		return
	}
	oldInstances, err := s.ListInstances()
	if err != nil {
		log.Errorf("Couldn't get all old instances -- %s", err.Error())
		return
	}
	delta := l.mergeInstances(allInstances, oldInstances)

	for _, oldPhoenixID := range delta.deadPhoenixIDs {
		if err := s.DeleteInstance(oldPhoenixID); err != nil {
			log.Errorf("Unable to delete instance %s from storage -- %s", oldPhoenixID, err.Error())
		}
	}

	for _, instance := range delta.instanceUpdates {
		if err := s.StoreInstance(instance); err != nil {
			log.Errorf("Unable to store instance %s to storage -- %s", instance, err.Error())
		}
	}
}
