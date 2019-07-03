package loop

import (
	"time"

	"github.com/theonlyjohnny/phoenix/internal/config"
	"github.com/theonlyjohnny/phoenix/internal/log"
	"github.com/theonlyjohnny/phoenix/pkg/backend"
	"github.com/theonlyjohnny/phoenix/pkg/storage"
)

type phoenixLoop struct {
	loopInterval time.Duration

	backend backend.Backend
	storage storage.Storage
}

// Start starts the main Phoenix loop
func Start(cfg *config.Config, s storage.Storage, b backend.Backend) error {

	loop, err := newPhoenixLoop(cfg, s, b)
	if err != nil {
		return err
	}
	log.Debugf("made loop: %#v", loop)
	loop.start()
	return nil
}

func newPhoenixLoop(cfg *config.Config, s storage.Storage, b backend.Backend) (*phoenixLoop, error) {
	return &phoenixLoop{
		cfg.LoopInterval,
		b,
		s,
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
	l.scaleClusters()
}

func (l *phoenixLoop) updateInstances() {
	s := l.storage
	allInstances := s.GetAllInstances()
	oldInstances := s.GetAllInstances()
	delta := l.mergeInstances(allInstances, oldInstances)

	for _, oldPhoenixID := range delta.deadPhoenixIDs {
		s.DeleteInstance(oldPhoenixID)
	}

	for _, instance := range delta.instanceUpdates {
		s.StoreInstance(instance)
	}
}
