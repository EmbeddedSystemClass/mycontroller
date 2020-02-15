package serial

import (
	m2s "github.com/mitchellh/mapstructure"
	s "github.com/tarm/serial"
)

// Config details
type Config struct {
	Portname string
	BaudRate int
}

// Client data
type Client struct {
	Config Config
	Port   *s.Port
}

// New serial client
func New(config map[string]string) (*Client, error) {
	var cfg Config
	err := m2s.Decode(config, &cfg)
	if err != nil {
		return nil, err
	}
	c := &s.Config{Name: cfg.Portname, Baud: cfg.BaudRate}
	port, err := s.OpenPort(c)
	if err != nil {
		return nil, err
	}
	d := &Client{
		Config: cfg,
		Port:   port,
	}
	return d, nil
}

func (d *Client) Write(data []byte) error {
	_, err := d.Port.Write(data)
	return err
}

// Close the driver
func (d *Client) Close() error {
	return d.Port.Close()
}
