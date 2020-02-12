package mongo

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	mg "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
func NewClient(config map[string]string) (*Client, error) {
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

func (c *Client) getCollection(entity string) *mg.Collection {
	return c.Client.Database(c.Config.Database).Collection(entity)
}

// Save date into database
func (c *Client) Save(entity string, data ...interface{}) error {
	if len(data) == 0 {
		return errors.New("No data provided")
	}
	cl := c.getCollection(entity)
	var err error
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
func (c *Client) Find(entity string, results interface{}) error {
	cl := c.getCollection(entity)
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
