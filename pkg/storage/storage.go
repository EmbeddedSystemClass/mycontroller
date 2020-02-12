package storage

import mongo "github.com/mycontroller-org/mycontroller/plugin/mongo"

// Client interface
type Client interface {
	Save(data ...interface{}) error
	Find(entityName string, results interface{}) error
}

// Entity returns entity name
type Entity interface {
	GetEntityName() string
}

// StorageClient to the world access
var StorageClient Client

// Init storage
func Init(config map[string]string) error {
	c, err := mongo.NewClient(config)
	if err != nil {
		return err
	}
	StorageClient = &c
	return nil
}
