package main

import (
	"fmt"
	"os"

	"github.com/akamensky/argparse"
	"github.com/theonlyjohnny/phoenix/internal/cloud"
	"github.com/theonlyjohnny/phoenix/internal/config"
	"github.com/theonlyjohnny/phoenix/internal/job"
	"github.com/theonlyjohnny/phoenix/internal/log"
	"github.com/theonlyjohnny/phoenix/internal/loop"
	"github.com/theonlyjohnny/phoenix/internal/server"
	"github.com/theonlyjohnny/phoenix/internal/storage"
)

func main() {
	parser := argparse.NewParser("phoenix", "PhoenixCli Entrypoint")
	s := parser.String("c", "config", &argparse.Options{Help: "path to config file", Default: "/etc/phoenix/config.json"})

	if err := parser.Parse(os.Args); err != nil {
		fmt.Print(parser.Usage(err))
		return
	}

	cfg := config.ReadConfigFromFs(*s)

	storage, err := storage.NewStorageEngine(cfg.StorageType)
	if err != nil {
		log.Errorf("unable to create storage -- exiting -- %s", err.Error())
		os.Exit(1)
	}

	cloud := cloud.NewCloudEngine(cfg, storage)

	manager, err := job.NewManager(storage, cloud)
	if err != nil {
		log.Errorf("unable to create job manager -- exiting -- %s", err.Error())
		os.Exit(1)
	}

	go func() {
		err := loop.Start(cfg, storage, cloud, manager)
		if err != nil {
			log.Errorf("unable to start loop -- exiting -- %s", err.Error())
			os.Exit(1)
		}
	}()
	server.Start(cfg, storage, manager)
}
