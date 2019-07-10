package loop

import (
	"time"

	"github.com/theonlyjohnny/phoenix/internal/config"
	"github.com/theonlyjohnny/phoenix/internal/job"
	"github.com/theonlyjohnny/phoenix/internal/log"
	"github.com/theonlyjohnny/phoenix/internal/storage"
	"github.com/theonlyjohnny/phoenix/pkg/backend"
)

type phoenixLoop struct {
	loopInterval time.Duration

	backend backend.Backend
	storage *storage.Engine
	manager *job.Manager
}

// Start starts the main Phoenix loop
func Start(cfg *config.Config, s *storage.Engine, b backend.Backend, m *job.Manager) error {

	loop, err := newPhoenixLoop(cfg, s, b, m)
	if err != nil {
		return err
	}
	loop.start()
	return nil
}

func newPhoenixLoop(cfg *config.Config, s *storage.Engine, b backend.Backend, m *job.Manager) (*phoenixLoop, error) {
	return &phoenixLoop{
		cfg.LoopInterval,
		b,
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
	b := l.backend
	allInstances := b.GetAllInstances()
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
