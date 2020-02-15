package client

// Client interface
type Client interface {
	Write() error
	Close() error
}
