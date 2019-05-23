package nullable

import (
	"database/sql"
	"encoding/json"
)

type Bool struct {
	sql.NullBool
}

func (b Bool) MarshalJSON() ([]byte, error) {
	if b.Valid {
		return json.Marshal(b.Bool)
	} else {
		return json.Marshal(nil)
	}
}

func (b *Bool) UnmarshalJSON(data []byte) error {
	var x *bool
	if err := json.Unmarshal(data, &x); err != nil {
		return err
	}
	if x != nil {
		b.Valid = true
		b.Bool = *x
	} else {
		b.Valid = false
	}
	return nil
}
