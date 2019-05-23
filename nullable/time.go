package nullable

import (
	"encoding/json"
	"github.com/lib/pq"
	"time"
)

type Time struct {
	pq.NullTime
}

func (t Time) MarshalJSON() ([]byte, error) {
	if t.Valid {
		return json.Marshal(t.Time)
	} else {
		return json.Marshal(nil)
	}
}

func (t *Time) UnmarshalJSON(data []byte) error {
	var x *time.Time
	if err := json.Unmarshal(data, &x); err != nil {
		return err
	}
	if x != nil {
		t.Valid = true
		t.Time = *x
	} else {
		t.Valid = false
	}
	return nil
}
