package loop

import (
	"github.com/theonlyjohnny/phoenix/internal/log"
)

func (l *phoenixLoop) scaleClusters() {
	s := l.storage
	clusters, err := s.ListClusters()
	if err != nil {
		log.Errorf("Unable to query clusters so unable to scale -- %s", err.Error())
	}

	for _, cluster := range clusters {
		log.Infof("Scaling %s", cluster)
	}
	//TODO actual scale logic
}
