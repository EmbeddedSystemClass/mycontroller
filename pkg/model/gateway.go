package model

import (
	mc "github.com/mycontroller-org/mycontroller/pkg"
	"github.com/mycontroller-org/mycontroller/pkg/storage"
	utils "github.com/mycontroller-org/mycontroller/pkg/utils"
)

// Gateway types
const (
	GwMySensors = "MY_SENSORS"
)

// Network types
const (
	NwMQTT   = "MQTT"
	NwSerial = "SERIAL"
)

// Acknowledgement data
type Acknowledgement struct {
	Enabled       bool   `json:"enabled"`
	StreamEnabled bool   `json:"streamEnabled"`
	RetryCount    bool   `json:"retryCount"`
	WaitTime      uint64 `json:"waitTime"`
}

// Gateway entity
type Gateway struct {
	Name        string          `json:"name"`
	UUID        string          `json:"uuid"`
	Enabled     bool            `json:"enabled"`
	Ack         Acknowledgement `json:"ack"`
	Type        string          `json:"type"`
	NetworkType string          `json:"networkType"`
	State       State           `json:"state"`
	Config      interface{}     `json:"config"`
}

// Reload gateway
func (g *Gateway) Reload() {
}

// Save gateway config into disk
func (g *Gateway) Save() error {
	if g.UUID == "" {
		g.UUID = utils.NewUUID()
	}
	return storage.StorageClient.Save(mc.EntGateway, g)
}

// SetState Updates state data
func (g *Gateway) SetState(s State) {
	g.State = s
	g.Save()
}

// GetEntityName returns the name of this entity
func (g *Gateway) GetEntityName() string {
	return mc.EntGateway
}

// GetUUID returns UUID of this object
func (g *Gateway) GetUUID() string {
	return g.UUID
}
