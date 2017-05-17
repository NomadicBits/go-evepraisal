package typedb

type TypeDB interface {
	GetType(typeName string) (EveType, bool)
	HasType(typeName string) bool
	Close() error
}

type EveType struct {
	ID              int64       `json:"id"`
	GroupID         int64       `json:"group_id"`
	Name            string      `json:"name"`
	Volume          float64     `json:"volume"`
	BaseComponenets []Component `json:"components"`
}

type Component struct {
	Quantity int64 `json:"quantity"`
	TypeID   int64 `json:"type_id"`
}