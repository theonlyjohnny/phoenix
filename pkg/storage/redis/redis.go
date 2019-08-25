package redis

import (
	"encoding/json"
	"fmt"

	"github.com/alicebob/miniredis"
	"github.com/go-redis/redis"
	"github.com/theonlyjohnny/phoenix/internal/config"
	"github.com/theonlyjohnny/phoenix/pkg/models"
	"github.com/theonlyjohnny/phoenix/pkg/storage"
)

type RedisStorage struct {
	clusterClient  *redis.Client
	instanceClient *redis.Client
}

func NewRedisStorage(cfg config.ComponentConfig) (storage.Storage, error) {

	clusterDBConfig := defaultConfig(0)
	instanceDBConfig := defaultConfig(1)

	if passedClusterCfg, err := cfg.GetNestedConfigComponent("cluster"); err == nil {
		clusterDBConfig = clusterDBConfig.Extend(passedClusterCfg)
	}

	if passedInstanceCfg, err := cfg.GetNestedConfigComponent("instance"); err == nil {
		instanceDBConfig = instanceDBConfig.Extend(passedInstanceCfg)
	}

	clusterClient, err := newClient(clusterDBConfig)
	if err != nil {
		return nil, err
	}

	instanceClient, err := newClient(instanceDBConfig)
	if err != nil {
		return nil, err
	}

	return &RedisStorage{
		clusterClient:  clusterClient,
		instanceClient: instanceClient,
	}, nil

}

func defaultConfig(db int) config.ComponentConfig {
	return config.ComponentConfig{
		"db":       db,
		"address":  "localhost:6379",
		"password": "",
	}
}

func newClient(cfg config.ComponentConfig) (*redis.Client, error) {

	address, err := cfg.GetStr("address")
	if err != nil {
		return nil, err
	}

	password, err := cfg.GetStr("password")
	if err != nil {
		return nil, err
	}

	db, err := cfg.GetInt("db")
	if err != nil {
		return nil, err
	}

	client := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
		DB:       db,
	})

	_, err = client.Ping().Result()
	if err != nil {
		return nil, err
	}

	return client, nil
}

func (r *RedisStorage) parseCluster(input interface{}) (*models.Cluster, error) {

	stringified, ok := input.(string)
	if !ok {
		return nil, fmt.Errorf("Input to parseCluster not a string")
	}

	var res models.Cluster

	err := json.Unmarshal([]byte(stringified), &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (r *RedisStorage) parseInstance(input interface{}) (*models.Instance, error) {

	stringified, ok := input.(string)
	if !ok {
		return nil, fmt.Errorf("Input to parseInstance not a string")
	}

	var res models.Instance

	err := json.Unmarshal([]byte(stringified), &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (r *RedisStorage) StoreInstance(key string, value *models.Instance) error {

	client := r.instanceClient

	stringified, err := json.Marshal(value)
	if err != nil {
		return err
	}

	return client.Set(key, stringified, 0).Err()
}

func (r *RedisStorage) StoreCluster(key string, value *models.Cluster) error {

	client := r.clusterClient

	stringified, err := json.Marshal(value)
	if err != nil {
		return err
	}

	return client.Set(key, stringified, 0).Err()
}

func (r *RedisStorage) ListInstances() (models.InstanceList, error) {
	client := r.instanceClient

	var cursor uint64
	allKeys := []string{}

	for {
		var keys []string
		var err error
		keys, cursor, err = client.Scan(cursor, "", 10).Result()
		if err != nil {
			return nil, err
		}
		allKeys = append(allKeys, keys...)
		if cursor == 0 {
			break
		}
	}

	vals, err := client.MGet(allKeys...).Result()
	if err != nil {
		return nil, err
	}

	result := make(models.InstanceList, len(vals), len(vals))

	for i, e := range vals {
		instance, err := r.parseInstance(e)
		if err != nil {
			return nil, err
		}
		result[i] = instance
	}

	return result, nil

}

func (r *RedisStorage) ListClusters() (models.ClusterList, error) {
	client := r.clusterClient

	var cursor uint64
	allKeys := []string{}

	for {
		var keys []string
		var err error
		keys, cursor, err = client.Scan(cursor, "", 10).Result()
		if err != nil {
			return nil, err
		}
		allKeys = append(allKeys, keys...)
		if cursor == 0 {
			break
		}
	}

	vals, err := client.MGet(allKeys...).Result()
	if err != nil {
		return nil, err
	}

	result := make(models.ClusterList, len(vals), len(vals))

	for i, e := range vals {
		cluster, err := r.parseCluster(e)
		if err != nil {
			return nil, err
		}
		result[i] = cluster
	}

	return result, nil

}

func (r *RedisStorage) GetInstance(key string) (*models.Instance, error) {
	client := r.instanceClient

	stringified, err := client.Get(key).Result()
	if err != nil {
		return nil, err
	}

	return r.parseInstance(stringified)
}

func (r *RedisStorage) GetCluster(key string) (*models.Cluster, error) {
	client := r.clusterClient

	stringified, err := client.Get(key).Result()
	if err != nil {
		return nil, err
	}

	return r.parseCluster(stringified)
}

func (r *RedisStorage) DeleteInstance(key string) error {
	client := r.instanceClient

	return client.Del(key).Err()
}

func (r *RedisStorage) DeleteCluster(key string) error {
	client := r.clusterClient

	return client.Del(key).Err()
}

func CreateTestDB() config.ComponentConfig {
	instanceDB, err := miniredis.Run()
	if err != nil {
		fmt.Printf("Unable to create miniredis for instances -- %s \n", err.Error())
	}

	clusterDB, err := miniredis.Run()
	if err != nil {
		fmt.Printf("Unable to create miniredis for clusters -- %s \n", err.Error())
	}

	return config.ComponentConfig{
		"instance": config.ComponentConfig{
			"db":       0,
			"address":  instanceDB.Addr(),
			"password": "",
		},
		"cluster": config.ComponentConfig{
			"db":       0,
			"address":  clusterDB.Addr(),
			"password": "",
		},
	}
}
