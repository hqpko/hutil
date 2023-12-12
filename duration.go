package hutils

import (
	"encoding/json"
	"errors"
	"time"
)

type Duration struct {
	time.Duration
}

func (d *Duration) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.String())
}

func (d *Duration) UnmarshalJSON(b []byte) (err error) {
	if len(b) < 2 {
		return errors.New("invalid duration")
	} else if b[0] != '"' || b[len(b)-1] != '"' {
		return errors.New("invalid duration")
	}
	d.Duration, err = time.ParseDuration(string(b[1 : len(b)-1]))
	return
}
