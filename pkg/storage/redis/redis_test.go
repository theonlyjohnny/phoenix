package redis

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/theonlyjohnny/phoenix/testsupport"
)

func newEmptyRedisStorage() *RedisStorage {
	return &RedisStorage{}
}

func TestParseCluster(t *testing.T) {
	cluster := testsupport.NewUniqueCluster()
	bytes, err := json.Marshal(cluster)

	assert.NoError(t, err)
	assert.NotEmpty(t, bytes)

	returnedCluster, err := newEmptyRedisStorage().parseCluster(string(bytes))

	assert.NoError(t, err)
	assert.Equal(t, cluster, returnedCluster)
}

func TestParseInstance(t *testing.T) {
	instance := testsupport.NewUniqueInstance()
	bytes, err := json.Marshal(instance)

	assert.NoError(t, err)
	assert.NotEmpty(t, bytes)

	returnedInstance, err := newEmptyRedisStorage().parseInstance(string(bytes))

	assert.True(t, instance.UpdatedDTTM.Equal(returnedInstance.UpdatedDTTM))
	returnedInstance.UpdatedDTTM = instance.UpdatedDTTM

	assert.NoError(t, err)
	assert.Equal(t, instance, returnedInstance)
}
