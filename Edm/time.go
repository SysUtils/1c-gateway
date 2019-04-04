package Edm

import (
	"encoding/json"
	"strings"
	"time"
)

type Time struct {
	time.Time
}

func (t Time) MarshalJSON() ([]byte, error) {
	val := t.Time.Format("2006-01-02T15:04:05")

	return json.Marshal(val)
}

func (t *Time) UnmarshalJSON(buf []byte) error {
	tt, err := time.Parse("2006-01-02T15:04:05", strings.Trim(string(buf), `"`))
	if err != nil {
		return err
	}
	t.Time = tt
	return nil
}

func (t *Time) Value() time.Time {
	return t.Time
}
