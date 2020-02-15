package mongo

import (
	"context"
	"errors"

	m2s "github.com/mitchellh/mapstructure"
	"github.com/mycontroller-org/mycontroller/pkg/interfaces"
	"go.mongodb.org/mongo-driver/bson"
	mg "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ctx = context.TODO()

// Config of the database
type Config struct {
	Name     string
	Database string
	URI      string
}

// Client of the mongo db
type Client struct {
	Client *mg.Client
	Config Config
}

// NewClient mongodb
func NewClient(config map[string]string) (*Client, error) {
	var cfg Config
	err := m2s.Decode(config, &cfg)
	if err != nil {
		return nil, err
	}
	clientOptions := options.Client().ApplyURI(cfg.URI)
	mc, err := mg.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}
	c := &Client{
		Config: cfg,
		Client: mc,
	}
	return c, nil
}

// Close the connection
func (c *Client) Close() error {
	return c.Client.Disconnect(ctx)
}

// Ping to the target database
func (c *Client) Ping() error {
	return c.Client.Ping(ctx, nil)
}

func (c *Client) getCollection(entity string) *mg.Collection {
	return c.Client.Database(c.Config.Database).Collection(entity)
}

// Save date into database
func (c *Client) Save(entity string, data interfaces.Entity) error {
	if data == nil {
		return errors.New("No data provided")
	}
	cl := c.getCollection(entity)
	// find the entity, if available update it
	or, err := cl.ReplaceOne(ctx, uuidFilter(data), data)
	if err != nil {
		return err
	}
	if or.MatchedCount == 0 {
		_, err := cl.InsertOne(ctx, data)
		if err != nil {
			return err
		}
	}

	return nil
}

// Find returns data
func (c *Client) Find(entity string, filter map[string]string, results interface{}) error {
	cl := c.getCollection(entity)
	bMap := bsonMap(filter)
	cur, err := cl.Find(ctx, bMap)
	if err != nil {
		return err
	}

	err = cur.All(ctx, results)
	if err != nil {
		return err
	}
	return nil
}

// FindOne returns data
func (c *Client) FindOne(entity string, filter map[string]string, result interface{}) error {
	cl := c.getCollection(entity)
	bMap := bsonMap(filter)
	res := cl.FindOne(ctx, bMap)
	if res.Err() != nil {
		return res.Err()
	}
	res.Decode(result)
	return nil
}

func bsonMap(gMap map[string]string) *bson.M {
	bMap := make(bson.M)
	for k, v := range gMap {
		bMap[k] = v
	}
	return &bMap
}

func uuidFilter(entity interfaces.Entity) *bson.M {
	return &bson.M{"uuid": entity.GetUUID()}
}
