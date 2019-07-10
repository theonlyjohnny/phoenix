package job

import (
	"github.com/theonlyjohnny/phoenix/internal/scale"
	"github.com/theonlyjohnny/phoenix/internal/storage"
	"github.com/theonlyjohnny/phoenix/pkg/backend"
)

//A Manager receives Events, recalculates state, and then applies any differences
type Manager struct {
	clusterLogic *scale.ClusterLogic
}

//NewManager returns a pointer to a newly instantiated Manager
func NewManager(storage *storage.Engine, backend backend.Backend) (*Manager, error) {
	return &Manager{
		scale.NewClusterLogic(storage, backend),
	}, nil
}

//TODO manage # of concurrect goroutines? <- limiting
//TODO make events go into queue and cancelable via context.Context
