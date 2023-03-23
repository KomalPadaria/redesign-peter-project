package nullable

import (
	"database/sql/driver"
	"encoding/json"

	"github.com/google/uuid"
)

type NullUUID struct {
	Valid bool
	UUID  uuid.UUID
}

func NewNullUUID(UUID uuid.UUID) *NullUUID {
	return &NullUUID{Valid: true, UUID: UUID}
}

// Scan implements the Scanner interface.
func (n *NullUUID) Scan(value interface{}) error {
	if value == nil {
		n.UUID, n.Valid = uuid.UUID{}, false
		return nil

	}
	n.Valid = true
	return n.UUID.Scan(value)

}

// Value implements the driver Valuer interface.
func (n NullUUID) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return n.UUID.Value()
}

func (n *NullUUID) MarshalJSON() ([]byte, error) {
	var val string
	if n.Valid {
		val = n.UUID.String()
	}

	return json.Marshal(val)
}
