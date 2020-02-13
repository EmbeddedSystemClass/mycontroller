package interfaces

// Entity returns entity name
type Entity interface {
	GetEntityName() string
	GetUUID() string
	Save() error
}
