package model

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
