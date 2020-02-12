package mongo

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	mg "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	s "github.com/mycontroller-org/mycontroller/pkg/storage"
)

var ctx = context.TODO()

// Config of the database
type Config struct {
	Database string
	URI      string
}

// Client of the mongo db
type Client struct {
	Client *mg.Client
	Config Config
}

// NewClient mongodb
func NewClient(config map[string]string) (s.Client, error) {
	c := &Client{
		Config: Config{
			Database: config["database"],
			URI:      config["uri"],
		},
	}
	clientOptions := options.Client().ApplyURI(c.Config.URI)
	cl, err := mg.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}
	c.Client = cl
	return c, nil
}

func (c *Client) getCollection(data interface{}) (*mg.Collection, error) {
	e, ok := data.(s.Entity)
	if !ok {
		return nil, errors.New("Provided data does not implemented Entity interface")
	}
	return c.getCollectionByName(e.GetEntityName()), nil
}

func (c *Client) getCollectionByName(name string) *mg.Collection {
	return c.Client.Database(c.Config.Database).Collection(name)
}

// Save date into database
func (c *Client) Save(data ...interface{}) error {
	if len(data) == 0 {
		return errors.New("No data provided")
	}
	cl, err := c.getCollection(&data[0])
	if err != nil {
		return err
	}
	if len(data) == 1 {
		_, err = cl.InsertOne(ctx, data[0])
	} else {
		_, err = cl.InsertMany(ctx, data)
	}

	if err != nil {
		return err
	}
	return nil
}

// Find returns data
func (c *Client) Find(entityName string, results interface{}) error {
	cl := c.getCollectionByName(entityName)
	cur, err := cl.Find(ctx, bson.M{})
	if err != nil {
		return err
	}

	err = cur.All(ctx, results)
	if err != nil {
		return err
	}
	return nil
}
