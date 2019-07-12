package job

import (
	"github.com/theonlyjohnny/phoenix/internal/cloud"
	logger "github.com/theonlyjohnny/phoenix/internal/log"
	"github.com/theonlyjohnny/phoenix/internal/scale"
	"github.com/theonlyjohnny/phoenix/internal/storage"
)

var log logger.Logger

func init() {
	log = logger.Log
}

//A Manager receives Events, recalculates state, and then applies any differences
type Manager struct {
	storage      *storage.Engine
	cloud        *cloud.Engine
	clusterLogic *scale.ClusterLogic
}

//NewManager returns a pointer to a newly instantiated Manager
func NewManager(storage *storage.Engine, cloud *cloud.Engine) (*Manager, error) {
	return &Manager{
		storage,
		cloud,
		scale.NewClusterLogic(storage, cloud),
	}, nil
}

//TODO manage # of concurrect goroutines? <- limiting
//TODO make events go into queue and cancelable via context.Context
