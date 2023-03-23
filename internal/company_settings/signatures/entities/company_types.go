package entities

import (
	"database/sql/driver"
	"encoding/json"
	"strings"

	"github.com/pkg/errors"
)

type CompanyTypes []string

// Value simply returns the JSON-encoded representation of the struct.
func (a CompanyTypes) Value() (driver.Value, error) {
	return json.Marshal(a)
}

// Scan makes the IndustryType implement the sql.Scanner interface. This method
// simply decodes a value into the string slice.
func (a *CompanyTypes) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &a)
}

// String simply returns the joined slice elements separated by comma.
func (a CompanyTypes) String() string {
	return strings.Join(a, ",")
}
