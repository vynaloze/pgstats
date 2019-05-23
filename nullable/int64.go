package nullable

import (
	"database/sql"
	"encoding/json"
)

type Int64 struct {
	sql.NullInt64
}

func (i Int64) MarshalJSON() ([]byte, error) {
	if i.Valid {
		return json.Marshal(i.Int64)
	} else {
		return json.Marshal(nil)
	}
}

func (i *Int64) UnmarshalJSON(data []byte) error {
	var x *int64
	if err := json.Unmarshal(data, &x); err != nil {
		return err
	}
	if x != nil {
		i.Valid = true
		i.Int64 = *x
	} else {
		i.Valid = false
	}
	return nil
}
