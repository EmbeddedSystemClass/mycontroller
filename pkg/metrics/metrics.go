package metrics

import (
	"github.com/mycontroller-org/mycontroller/plugin/metrics/influx"
)

// Client interface
type Client interface {
}

// MetricsClient to the world access
var MetricsClient Client

// Init storage
func Init(config map[string]string) error {
	c, err := influx.NewClient(config)
	if err != nil {
		return err
	}
	MetricsClient = c
	return nil
}
