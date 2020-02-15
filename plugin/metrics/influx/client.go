package influx

import (
	"context"

	"github.com/influxdata/influxdb-client-go"
	m2s "github.com/mitchellh/mapstructure"
)

var ctx = context.TODO()

// Config of the database
type Config struct {
	Name     string
	Database string
	URI      string
	Token    string
	Username string
	Password string
}

// Client of the influxdb
type Client struct {
	Client *influxdb.Client
	Config Config
}

// NewClient of influxdb
func NewClient(config map[string]string) (*Client, error) {
	var cfg Config
	err := m2s.Decode(config, &cfg)
	if err != nil {
		return nil, err
	}
	var ic *influxdb.Client
	if cfg.Token != "" {
		ic, err = influxdb.New(cfg.URI, cfg.Token)
	} else {
		ic, err = influxdb.New(cfg.URI, "", influxdb.WithUserAndPass(cfg.Username, cfg.Password))
	}
	if err != nil {
		return nil, err
	}
	c := &Client{
		Config: cfg,
		Client: ic,
	}
	return c, nil
}

// Ping to target database
func (c *Client) Ping() error {
	return c.Client.Ping(ctx)
}

// Close the influxdb connection
func (c *Client) Close() error {
	return c.Client.Close()
}
