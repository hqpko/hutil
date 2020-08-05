package hutils

import (
	"testing"
	"time"
)

func TestWaitTimout(t *testing.T) {
	wt := NewWaitTimeout()
	wt.Add(1)
	go func() {
		wt.Done()
	}()
	if success := <-wt.Wait(100 * time.Millisecond); !success { // true
		t.Errorf("waitTimeout.Wait fail")
	}

	wt = NewWaitTimeout()
	wt.Add(1)
	go func() {
		time.Sleep(200 * time.Millisecond)
		wt.Done()
	}()
	if success := <-wt.Wait(100 * time.Millisecond); success { // false
		t.Errorf("waitTimeout.Wait fail")
	}
}
