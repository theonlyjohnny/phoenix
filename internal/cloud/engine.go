package cloud

import (
	"fmt"

	"github.com/theonlyjohnny/phoenix/internal/config"
	logger "github.com/theonlyjohnny/phoenix/internal/log"
	"github.com/theonlyjohnny/phoenix/internal/storage"
	"github.com/theonlyjohnny/phoenix/pkg/cloud"
	"github.com/theonlyjohnny/phoenix/pkg/cloud/ec2"
)

var log logger.Logger

func init() {
	log = logger.Log
}

type Engine struct {
	baseCfgs      map[string]config.ComponentConfig
	providerCache map[string]cloud.Provider

	storage *storage.Engine
}

func NewCloudEngine(cfg *config.Config, s *storage.Engine) *Engine {
	return &Engine{
		cfg.CloudProviderConfig,
		map[string]cloud.Provider{},
		s,
	}
}

func (e *Engine) GetCloudProvider(clusterName string, cfg *config.ComponentConfig) (cloud.Provider, error) {
	// log.Debugf("getCloudProvider cache: %#v name: %s base: %s", e.providerCache, clusterName, e.baseCfgs)
	if provider, ok := e.providerCache[clusterName]; ok {
		return provider, nil
	}

	var finalCfg config.ComponentConfig

	var provider cloud.Provider
	var providerType string
	var err error

	if cluster, err := e.storage.GetCluster(clusterName); err == nil {
		providerType = cluster.CloudProviderType
	} else {
		return provider, fmt.Errorf("Failed to create cloud provider for %s -- %s", clusterName, err.Error())
	}

	if baseCfg, ok := e.baseCfgs[providerType]; err == nil && ok {
		finalCfg = baseCfg
	}

	if cfg != nil {
		finalCfg = finalCfg.Extend(*cfg)
	}

	switch providerType {
	case "ec2":
		provider, err = ec2.NewEC2CloudProvider(finalCfg)
	default:
		log.Errorf("Unable to find cloud provider with type %s", providerType)
		return provider, fmt.Errorf("unknown provider: %s", providerType)
	}

	if err == nil {
		e.providerCache[clusterName] = provider
	}

	return provider, err
}
