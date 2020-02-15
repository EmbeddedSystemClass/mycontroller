package mqtt

import (
	"time"

	paho "github.com/eclipse/paho.mqtt.golang"
	m2s "github.com/mitchellh/mapstructure"
)

// Config details
type Config struct {
	URL      string
	Username string
	Password string
}

// Client data
type Client struct {
	Config Config
	Client paho.Client
}

// New mqtt driver
func New(config map[string]string) (*Client, error) {
	var cfg Config
	err := m2s.Decode(config, &cfg)
	if err != nil {
		return nil, err
	}

	opts := paho.NewClientOptions()
	opts.AddBroker(cfg.URL)
	opts.SetUsername(cfg.Username)
	opts.SetPassword(cfg.Password)
	opts.SetClientID("")

	c := paho.NewClient(opts)
	token := c.Connect()
	for !token.WaitTimeout(3 * time.Second) {
	}
	if err := token.Error(); err != nil {
		return nil, err
	}

	d := &Client{
		Config: cfg,
		Client: c,
	}
	return d, nil
}

// Write publishes a payload
func (d *Client) Write(data []byte) error {
	token := d.Client.Publish("topic", 0, false, "payload")
	return token.Error()
}

// Close the driver
func (d *Client) Close() error {
	if d.Client.IsConnected() {
		d.Client.Disconnect(0)
	}
	return nil
}
