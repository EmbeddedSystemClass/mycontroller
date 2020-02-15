package ethernet

import (
	"net"
	"net/url"

	m2s "github.com/mitchellh/mapstructure"
)

// Config details
type Config struct {
	URI string
}

// Client data
type Client struct {
	Config Config
	Client net.Conn
}

// New ethernet driver
func New(config map[string]string) (*Client, error) {
	var cfg Config
	err := m2s.Decode(config, &cfg)
	if err != nil {
		return nil, err
	}

	uri, err := url.Parse(cfg.URI)
	if err != nil {
		return nil, err
	}

	c, err := net.Dial(uri.Scheme, uri.Host)
	if err != nil {
		return nil, err
	}

	d := &Client{
		Config: cfg,
		Client: c,
	}
	return d, nil
}

// Write sends a payload
func (d *Client) Write(data []byte) error {
	_, err := d.Client.Write(data)
	return err
}

// Close the connection
func (d *Client) Close() error {
	return d.Close()
}
