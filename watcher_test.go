package hutils

import (
	"testing"
	"time"
)

func TestWatcher(t *testing.T) {
	w := NewWatcher()
	go w.Notify()
	if !w.Watch(100 * time.Millisecond) {
		t.Errorf("watcher.watch fail")
	}

	go func() {
		time.Sleep(200 * time.Millisecond)
		w.Notify()
	}()
	if w.Watch(100 * time.Millisecond) {
		t.Errorf("watcher.watch fail")
	}
}
