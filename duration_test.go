package hutils

import (
	"encoding/json"
	"testing"
	"time"
)

type d struct {
	D Duration `json:"d"`
}

func TestDuration(t *testing.T) {
	s := `{"d": "1h"}`
	d := d{}
	err := json.Unmarshal([]byte(s), &d)
	if err != nil {
		t.Error(err)
	}
	if d.D.Duration != time.Hour {
		t.Error("error")
	}
}
