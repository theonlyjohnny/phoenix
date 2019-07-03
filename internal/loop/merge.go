package loop

import (
	"github.com/theonlyjohnny/phoenix/internal/cluster"
	"github.com/theonlyjohnny/phoenix/internal/instance"
	"github.com/theonlyjohnny/phoenix/internal/log"
)

type mergedInstancesDelta struct {
	instanceUpdates []*instance.Instance
	deadPhoenixIDs  []string
}

func (l *phoenixLoop) mergeInstances(allInstances, oldInstances []*instance.Instance) *mergedInstancesDelta {
	s := l.storage
	clusters := s.ListClusters()

	updatedInstances := []*instance.Instance{}
	deadPhoenixIDs := []string{}

	relevantInstancesIDMap := map[string]*instance.Instance{}
	oldPhoenixIDs := map[string]*instance.Instance{}
	allPhoenixIDs := map[string]bool{}

	for _, oldInstance := range oldInstances {
		oldPhoenixIDs[oldInstance.PhoenixID] = oldInstance
	}

	for _, allInstance := range allInstances {
		allPhoenixIDs[allInstance.PhoenixID] = false
	}

	log.Debugf("merging all: %v and old: %v", allInstances, oldInstances)

	for _, instance := range allInstances {
		var instanceCluster *cluster.Cluster

		for _, potentialCluster := range clusters {
			if potentialCluster.HasInstance(instance) {
				instanceCluster = potentialCluster
				break
			}
		}

		if instanceCluster == nil {
			if _, ok := oldPhoenixIDs[instance.PhoenixID]; ok {
				log.Warnf("Found a managed instance with no known cluster, did a cluster get deleted? %s", instance)
				deadPhoenixIDs = append(deadPhoenixIDs, instance.PhoenixID)
			} else {
				log.Debugf("Found an instance with no known cluster -- %s", instance)
			}
			continue
		}

		relevantInstancesIDMap[instance.PhoenixID] = instance

		if oldInstance, ok := oldPhoenixIDs[instance.PhoenixID]; !ok {
			log.Infof("New Instance! %s", instance)
			updatedInstances = append(updatedInstances, instance)
		} else if oldInstance != instance {
			updatedInstances = append(updatedInstances, instance)
		}
	}

	for _, instance := range oldInstances {
		if _, ok := allPhoenixIDs[instance.PhoenixID]; !ok {
			//TODO check if in creating state and skip?
			log.Warnf("Found an instance not reported by backend -- %s", instance)
			deadPhoenixIDs = append(deadPhoenixIDs, instance.PhoenixID)
		}
	}

	return &mergedInstancesDelta{
		updatedInstances,
		deadPhoenixIDs,
	}
}