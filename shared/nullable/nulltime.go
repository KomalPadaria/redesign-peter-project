package nullable

import (
	"database/sql"
	"encoding/json"
	"strings"
	"time"
)

type NullTime struct {
	sql.NullTime
}

func NewNullTime(time time.Time) NullTime {
	return NullTime{sql.NullTime{
		Time:  time,
		Valid: true,
	}}
}

func (n *NullTime) MarshalJSON() ([]byte, error) {
	var val string
	if n.Valid {
		val = n.Time.Format(time.RFC3339)
	}

	return json.Marshal(val)
}

func (nt *NullTime) UnmarshalJSON(data []byte) error {
	nt.Valid = false
	s := strings.Trim(string(data), "\"")
	if s == "null" {
		return nil
	}
	if s != "" {
		t, err := time.Parse(time.RFC3339, s)
		if err != nil {
			return err
		}
		if !t.IsZero() {
			nt.Valid = true
			nt.Time = t
		}
	}

	return nil
}
