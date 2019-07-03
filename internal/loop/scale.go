package loop

import (
	"github.com/theonlyjohnny/phoenix/internal/log"
)

func (l *phoenixLoop) scaleClusters() {
	s := l.storage
	clusters := s.ListClusters()
	log.Debugf("storage: %#v", l.storage)

	for _, cluster := range clusters {
		log.Infof("Scaling %s", cluster)
	}
}
