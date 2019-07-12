package log

import (
	"fmt"
	"os"

	"github.com/theonlyjohnny/jogger-go/logger"
)

type Logger interface {
	Debugf(string, ...interface{})
	Infof(string, ...interface{})
	Errorf(string, ...interface{})
	Warnf(string, ...interface{})
}

var Log Logger

func setupLogger() error {
	opts := logger.Config{
		AppName:    "gorgon",
		LogLevel:   "debug",
		LogConsole: true,
		LogSyslog:  nil,
	}
	var loggerErr error
	Log, loggerErr = logger.CreateLogger(opts)
	return loggerErr
}

func init() {
	if err := setupLogger(); err != nil {
		fmt.Printf("Failed to setup logger ? %s \n", err)
		os.Exit(1)
	}
}
