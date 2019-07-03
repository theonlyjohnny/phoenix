package log

import (
	"fmt"
	"os"

	"github.com/theonlyjohnny/jogger-go/logger"
)

var log *logger.Logger

func setupLogger() error {
	opts := logger.Config{
		AppName:    "gorgon",
		LogLevel:   "debug",
		LogConsole: true,
		LogSyslog:  nil,
	}
	var loggerErr error
	log, loggerErr = logger.CreateLogger(opts)
	return loggerErr
}

func init() {
	if err := setupLogger(); err != nil {
		fmt.Printf("Failed to setup logger ? %s \n", err)
		os.Exit(1)
	}
}

//Debugf prints an formated debug log
func Debugf(msg string, args ...interface{}) {
	log.Debugf(msg, args...)
}

//Infof prints an formated info log
func Infof(msg string, args ...interface{}) {
	log.Infof(msg, args...)
}

//Errorf prints an formated error log
func Errorf(msg string, args ...interface{}) {
	log.Errorf(msg, args...)
}

//Warnf prints an formated warning
func Warnf(msg string, args ...interface{}) {
	log.Warnf(msg, args...)
}
