package storage

import (
	"github.com/mycontroller-org/mycontroller/pkg/interfaces"
	"github.com/mycontroller-org/mycontroller/plugin/storage/mongo"
)

// Client interface
type Client interface {
	Save(entity string, data interfaces.Entity) error
	Find(entityName string, filter map[string]string, results interface{}) error
	FindOne(entityName string, filter map[string]string, result interface{}) error
}

// StorageClient to the world access
var StorageClient Client

// Init storage
func Init(config map[string]string) error {
	c, err := mongo.NewClient(config)
	if err != nil {
		return err
	}
	StorageClient = c
	return nil
}
