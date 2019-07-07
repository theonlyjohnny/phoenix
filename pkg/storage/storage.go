package storage

//EntityType is used to distinguish what kind of entity is being stored
type EntityType string

const (
	InstanceEntityType = EntityType("Instance")
	ClusterEntityType  = EntityType("Cluster")
)

//Storage stores interfaces to some (ideally remote) persistent storage
type Storage interface {
	Store(EntityType, string, interface{}) error
	List(EntityType) ([]interface{}, error)
	Get(EntityType, string) (interface{}, error)
	Delete(EntityType, string) error
}
