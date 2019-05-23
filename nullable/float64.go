package nullable

import (
	"database/sql"
	"encoding/json"
)

type Float64 struct {
	sql.NullFloat64
}

func (f Float64) MarshalJSON() ([]byte, error) {
	if f.Valid {
		return json.Marshal(f.Float64)
	} else {
		return json.Marshal(nil)
	}
}

func (f *Float64) UnmarshalJSON(data []byte) error {
	var x *float64
	if err := json.Unmarshal(data, &x); err != nil {
		return err
	}
	if x != nil {
		f.Valid = true
		f.Float64 = *x
	} else {
		f.Valid = false
	}
	return nil
}
