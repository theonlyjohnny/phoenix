package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/theonlyjohnny/phoenix/internal/config"
	"github.com/theonlyjohnny/phoenix/internal/log"
	"github.com/theonlyjohnny/phoenix/pkg/backend"
	"github.com/theonlyjohnny/phoenix/pkg/storage"
)

//StorageKey is the key to access the instantiated storage.Storage within a request context
const StorageKey = string(iota)

// Start starts the HTTP server on the specified port
func Start(cfg *config.Config, s *storage.Storage, b backend.Backend) error {
	port := cfg.Port

	log.Infof("Starting server on %d \n", port)

	r := gin.New()
	r.Use(gin.Recovery())

	api := r.Group("/api")
	api.Use(authMiddleware())
	api.Use(logMiddleware())
	api.Use(storageMiddleware(s))

	cluster := api.Group("/cluster")
	cluster.POST("/", postClusterHandler)

	r.Run(fmt.Sprintf(":%d", port))
	return nil
}
