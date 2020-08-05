package hutils

import (
	"testing"
	"time"
)

func TestWatcher(t *testing.T) {
	w := NewWatcher()
	go w.Notify()
	if success := <-w.Watch(100 * time.Millisecond); !success {
		t.Errorf("watcher.watch fail")
	}

	go func() {
		time.Sleep(200 * time.Millisecond)
		w.Notify()
	}()
	if success := <-w.Watch(100 * time.Millisecond); success {
		t.Errorf("watcher.watch fail")
	}
}
