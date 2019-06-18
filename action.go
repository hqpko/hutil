package hutils

import (
	"sync"

	"github.com/hqpko/hconcurrent"
)

type action struct {
	actionType int32
	value      interface{}
}

type actionPool struct {
	pool *sync.Pool
}

func newActionPool() *actionPool {
	return &actionPool{pool: &sync.Pool{New: func() interface{} {
		return &action{}
	}}}
}

func (ap *actionPool) get(actionType int32, value interface{}) *action {
	action := ap.pool.Get().(*action)
	action.actionType = actionType
	action.value = value
	return action
}

func (ap *actionPool) put(action *action) {
	action.actionType = 0
	action.value = nil
	ap.pool.Put(action)
}

type ActionChannel struct {
	lock        *sync.RWMutex
	mainChannel *hconcurrent.Concurrent
	pool        *actionPool
	actions     map[int32]func(interface{})
}

func NewActionChannel() *ActionChannel {
	return NewActionChannelWithOption(1<<10, 1)
}

func NewActionChannelWithOption(channelSize, goroutineCount int) *ActionChannel {
	ac := &ActionChannel{lock: &sync.RWMutex{}, pool: newActionPool(), actions: map[int32]func(interface{}){}}
	ac.mainChannel = hconcurrent.NewConcurrent(channelSize, goroutineCount, ac.doAction)
	ac.mainChannel.Start()
	return ac
}

func (ac *ActionChannel) Register(actionType int32, handler func(value interface{})) {
	ac.lock.Lock()
	defer ac.lock.Unlock()
	ac.actions[actionType] = handler
}

func (ac *ActionChannel) Input(actionType int32, value interface{}) bool {
	return ac.mainChannel.Input(ac.pool.get(actionType, value))
}

func (ac *ActionChannel) MustInput(actionType int32, value interface{}) {
	ac.mainChannel.MustInput(ac.pool.get(actionType, value))
}

func (ac *ActionChannel) Stop() {
	ac.mainChannel.Stop()
}

func (ac *ActionChannel) doAction(i interface{}) interface{} {
	ac.lock.RLock()
	defer ac.lock.RUnlock()
	if action, ok := i.(*action); ok {
		if f, ok := ac.actions[action.actionType]; ok {
			f(action.value)
		}
		ac.pool.put(action)
	}
	return nil
}
