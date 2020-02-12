package model

import utils "github.com/mycontroller-org/mycontroller/pkg/utils"
import mc "github.com/mycontroller-org/mycontroller/pkg"

// Gateway types
const (
	GwMySensors = "MY_SENSORS"
)

// Network types
const (
	NwMQTT   = "MQTT"
	NwSerial = "SERIAL"
)

// State
const (
	StateUp          = "UP"
	StateDown        = "DOWN"
	StateUnavailable = "UNAVAILABLE"
)

// State data
type State struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Since   uint64 `json:"since"`
}

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
func (g *Gateway) Save() {
	if g.UUID == "" {
		g.UUID = utils.NewUUID()
	}
	// add code to save
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
