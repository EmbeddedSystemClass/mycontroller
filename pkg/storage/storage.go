package storage

import (
	"github.com/mycontroller-org/mycontroller/plugin/storage/mongo"
)

// Client interface
type Client interface {
	Save(entity string, data ...interface{}) error
	Find(entityName string, results interface{}) error
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
