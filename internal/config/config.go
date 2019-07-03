package config

import (
	"encoding/json"
	"io/ioutil"
	"time"

	"github.com/theonlyjohnny/phoenix/internal/log"
)

var (
	validBackends = []string{"ec2"}
	validStorages = []string{"local"}
)

//BackendConfig is an arbitrary JSON interface for use by individual backends
type BackendConfig map[string]interface{}

//StorageConfig is an arbitrary JSON interface for use by individual backends
// type StorageConfig map[string]interface{}

//Config stores all the config values read from a config file
type Config struct {
	Port         int           `json:"port"`
	LoopInterval time.Duration `json:"loop_interval_ns"`

	BackendType   string        `json:"backend_type"`
	BackendConfig BackendConfig `json:"backend_config"`

	StorageType string `json:"storage_type"`
	// StorageConfig StorageConfig `json:"storage_config"`
}

func defaultConfig() *Config {
	return &Config{
		Port:          9000,
		LoopInterval:  time.Second * 10,
		BackendType:   "ec2",
		BackendConfig: BackendConfig{},
		StorageType:   "local",
	}
}

//ReadConfigFromFs reads from the specified path and merges any options onto a default instance of Config
func ReadConfigFromFs(path string) *Config {
	log.Debugf("searching for config file @ %s", path)
	cfg := defaultConfig()
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Warnf("Unable to read file, using default config -- %s", err.Error())
		return cfg
	}

	if err := json.Unmarshal(data, cfg); err != nil {
		log.Warnf("Unable to combine config with default config -- %s", err.Error())
	}

	isValidBackend := strContains(cfg.BackendType, validBackends)
	isValidStorage := strContains(cfg.StorageType, validStorages)

	if !isValidBackend {
		defaultBackend := defaultConfig().BackendType
		log.Warnf("invalid backend specified %s -- falling back to %s", cfg.BackendType, defaultBackend)
		cfg.BackendType = defaultBackend
	}

	if !isValidStorage {
		defaultStorage := defaultConfig().StorageType
		log.Warnf("invalid storage specified %s -- falling back to %s", cfg.StorageType, defaultStorage)
		cfg.StorageType = defaultStorage
	}

	return cfg
}

func strContains(search string, in []string) bool {
	for _, e := range in {
		if e == search {
			return true
		}
	}
	return false
}
