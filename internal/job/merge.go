package job

import "github.com/theonlyjohnny/phoenix/pkg/models"

type mergedInstancesDelta struct {
	instanceUpdates models.InstanceList
	deadPhoenixIDs  []string
}

func (m *Manager) mergeInstances(allInstances, oldInstances models.InstanceList) *mergedInstancesDelta {
	s := m.storage
	clusters, err := s.ListClusters()
	if err != nil {
		log.Errorf("unable to list clusters and thus unable to merge instances -- %s", err.Error())
	}

	updatedInstances := models.InstanceList{}
	deadPhoenixIDs := []string{}

	relevantInstancesIDMap := map[string]*models.Instance{}
	oldPhoenixIDs := map[string]*models.Instance{}
	allPhoenixIDs := map[string]bool{}

	for _, oldInstance := range oldInstances {
		oldPhoenixIDs[oldInstance.PhoenixID] = oldInstance
	}

	for _, allInstance := range allInstances {
		allPhoenixIDs[allInstance.PhoenixID] = false
	}

	log.Debugf("merging all: %v and old: %v", allInstances, oldInstances)

	for _, instance := range allInstances {
		var instanceCluster *models.Cluster

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
				log.Debugf("Found a new instance with no known cluster -- %s", instance)
				updatedInstances = append(updatedInstances, instance)
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
			log.Warnf("Found an instance not reported by backend -- %s", instance)
			deadPhoenixIDs = append(deadPhoenixIDs, instance.PhoenixID)
		}
	}

	return &mergedInstancesDelta{
		updatedInstances,
		deadPhoenixIDs,
	}
}
