package nullable

import (
	"database/sql"
	"encoding/json"
)

type String struct {
	sql.NullString
}

func (i String) MarshalJSON() ([]byte, error) {
	if i.Valid {
		return json.Marshal(i.String)
	} else {
		return json.Marshal(nil)
	}
}

func (i *String) UnmarshalJSON(data []byte) error {
	var x *string
	if err := json.Unmarshal(data, &x); err != nil {
		return err
	}
	if x != nil {
		i.Valid = true
		i.String = *x
	} else {
		i.Valid = false
	}
	return nil
}
