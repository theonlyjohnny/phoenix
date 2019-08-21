package main

import (
	"net/url"
	"os"

	loop "github.com/theonlyjohnny/phoenix/internal/clientloop"
	logger "github.com/theonlyjohnny/phoenix/internal/log"
)

func main() {
	log := logger.Log

	phoenixID := os.Getenv("PHOENIX_ID")
	serverLocation := os.Getenv("PHOENIX_MASTER_LOCATION")

	if phoenixID == "" {
		log.Errorf("Env var PHOENIX_ID is required")
		os.Exit(1)
	}

	if serverLocation == "" {
		log.Errorf("Env var PHOENIX_MASTER_LOCATION is required")
		os.Exit(1)
	}

	url, err := url.Parse(serverLocation)
	if err != nil {
		log.Errorf("Unable to parse PHOENIX_MASTER_LOCATION into valid URL -- %s", err.Error())
		os.Exit(1)
	}

	loop.Start(phoenixID, url)
}
