package nullable

import (
	"database/sql"
	"encoding/json"
)

type NullString struct {
	sql.NullString
}

func NewNullString(s string) NullString {
	return NullString{sql.NullString{
		String: s,
		Valid:  true,
	}}
}

func (ns *NullString) MarshalJSON() ([]byte, error) {
	var val string
	if ns.Valid {
		val = ns.String
	}

	return json.Marshal(val)
}

func (ns *NullString) UnmarshalJSON(data []byte) error {
	var s *string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	if s != nil {
		ns.Valid = true
		ns.String = *s
	} else {
		ns.Valid = false
	}
	return nil
}
