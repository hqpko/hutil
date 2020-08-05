package hutils

import (
	"sync/atomic"
	"time"
)

// WaitTimeout 带有超时的 WaitGroup，避免了 wg.Wait() 可能引起的 blocked forever，协程永远无法被释放
type WaitTimeout struct {
	done  chan struct{}
	count int32
}

// NewWaitTimeout 必须使用 NewWaitTimeout 实例化 WaitTimeout
func NewWaitTimeout() *WaitTimeout {
	return &WaitTimeout{done: make(chan struct{})}
}

// Wait, if timeout, return false
func (wt *WaitTimeout) Wait(timeout time.Duration) bool {
	if count := atomic.LoadInt32(&wt.count); count <= 0 {
		return true
	}

	select {
	case <-wt.done:
		return true
	case <-time.After(timeout):
		return false
	}
}

func (wt *WaitTimeout) Add(n int32) *WaitTimeout {
	atomic.AddInt32(&wt.count, n)
	return wt
}

func (wt *WaitTimeout) Done() {
	if count := atomic.AddInt32(&wt.count, -1); count <= 0 {
		select {
		case wt.done <- struct{}{}:
		default:
		}
	}
}

func WaitTimeoutFunc(timeout time.Duration, f func()) bool {
	wt := NewWaitTimeout().Add(1)
	go func() {
		f()
		wt.Done()
	}()
	return wt.Wait(timeout)
}
