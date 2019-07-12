package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/theonlyjohnny/phoenix/internal/config"
	"github.com/theonlyjohnny/phoenix/internal/job"
	logger "github.com/theonlyjohnny/phoenix/internal/log"
	"github.com/theonlyjohnny/phoenix/internal/storage"
)

var log logger.Logger

const (
	wrapperKey = string(iota)
)

func init() {
	log = logger.Log
}

// Start starts the HTTP server on the specified port
func Start(cfg *config.Config, s *storage.Engine, m *job.Manager) error {
	port := cfg.Port

	log.Infof("Starting server on %d \n", port)

	r := gin.New()
	r.Use(gin.Recovery())

	api := r.Group("/api")
	api.Use(authMiddleware())
	api.Use(logMiddleware())
	api.Use(storageMiddleware(s, m))

	cluster := api.Group("/cluster")
	cluster.POST("/", unWrapHandler(postClusterHandler))

	r.Run(fmt.Sprintf(":%d", port))
	return nil
}
