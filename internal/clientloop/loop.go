package clientloop

import (
	"net/http"
	"net/url"
	"time"

	logger "github.com/theonlyjohnny/phoenix/internal/log"
	"github.com/theonlyjohnny/phoenix/pkg/models"
)

var log logger.Logger

func init() {
	log = logger.Log
}

const (
	loopInterval = time.Second * 5
)

type clientLooper struct {
	ticker *time.Ticker
	http   *http.Client

	phoenixID      string
	serverLocation *url.URL
}

func Start(phoenixID string, serverLocation *url.URL) {

	client := getNewHTTPClient()

	clientLoop := clientLooper{
		time.NewTicker(loopInterval),
		client,
		phoenixID,
		serverLocation,
	}

	clientLoop.start()

}

func (l *clientLooper) start() {
	for range l.ticker.C {
		l.loop()
	}
}

func (l *clientLooper) loop() {
	status, err := l.determineStatus()
	if err != nil {
		log.Errorf("Unable to determine status -- %s", err.Error())
		return
	}
	err = l.sendStatus(status)
	if err != nil {
		log.Errorf("Unable to send status -- %s", err.Error())
		return
	}
	log.Debugf("Successfully sent status to server")
}

func (l *clientLooper) determineStatus() (models.Status, error) {
	//TODO
	return models.Status{
		CPUUsage: 50.0,
		MemUsage: 10.0,
		Healthy:  true,
	}, nil
}

func (l *clientLooper) sendStatus(status models.Status) error {
	return l.postHTTP("status", status)
}
