package loop

import (
	"github.com/theonlyjohnny/phoenix/internal/log"
)

func (l *phoenixLoop) scaleClusters() {
	s := l.storage
	clusters := s.ListClusters()

	for _, cluster := range clusters {
		log.Infof("Scaling %s", cluster)
	}
	//TODO actual scale logic
}
