package main

import (
	"fmt"
	"os"

	"github.com/akamensky/argparse"
	"github.com/theonlyjohnny/phoenix/internal/config"
	"github.com/theonlyjohnny/phoenix/internal/job"
	"github.com/theonlyjohnny/phoenix/internal/log"
	"github.com/theonlyjohnny/phoenix/internal/loop"
	"github.com/theonlyjohnny/phoenix/internal/server"
	"github.com/theonlyjohnny/phoenix/pkg/backend"
	"github.com/theonlyjohnny/phoenix/pkg/storage"
)

func main() {
	parser := argparse.NewParser("phoenix", "PhoenixCli Entrypoint")
	s := parser.String("c", "config", &argparse.Options{Help: "path to config file", Default: "/etc/phoenix/config.json"})

	if err := parser.Parse(os.Args); err != nil {
		fmt.Print(parser.Usage(err))
		return
	}

	cfg := config.ReadConfigFromFs(*s)
	log.Debugf("cfg: %+v, backendCfg: %+v", cfg, cfg.BackendConfig)

	backend, err := backend.GetBackendByType(cfg.BackendType, cfg.BackendConfig)
	if err != nil {
		log.Errorf("unable to create backend -- exiting -- %s", err.Error())
		os.Exit(1)
	}

	storage, err := storage.GetStorageByType(cfg.StorageType)
	if err != nil {
		log.Errorf("unable to create storage -- exiting -- %s", err.Error())
		os.Exit(1)
	}

	manager, err := job.NewManager()
	if err != nil {
		log.Errorf("unable to create job manager -- exiting -- %s", err.Error())
		os.Exit(1)
	}

	go func() {
		err := loop.Start(cfg, *storage, backend)
		if err != nil {
			log.Errorf("unable to start loop -- exiting -- %s", err.Error())
			os.Exit(1)
		}
	}()
	server.Start(cfg, storage, backend, manager)
}
