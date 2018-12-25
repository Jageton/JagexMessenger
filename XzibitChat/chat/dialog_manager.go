package chat

import (
	"errors"
	"sync"
)

type DialogManager struct {
	hubs   map[int64]*Dialog
	mutex *sync.RWMutex
}

func NewHubManager() *DialogManager {
	return &DialogManager{
		hubs : make(map[int64]*Dialog),
		mutex : &sync.RWMutex{},
	}
}

func (h *DialogManager) Add(id int64, hub *Dialog) error {
	if h.Contains(id) {
		return errors.New("user is already exists")
	}
	h.mutex.Lock()
	defer h.mutex.Unlock()
	h.hubs[id] = hub
	return nil
}

func (h *DialogManager) AddEmpty(id int64) error {
	if h.Contains(id) {
		return errors.New("user is already exists")
	}
	h.mutex.Lock()
	defer h.mutex.Unlock()
	h.hubs[id] = NewHub()
	return nil
}

func (h *DialogManager) Remove(id int64) bool {
	h.mutex.Lock()
	defer h.mutex.Unlock()
	var ok bool
	if _, ok = h.hubs[id]; ok {
		delete(h.hubs, id)
	}
	return ok
}

func (h *DialogManager) Get(id int64) (*Dialog, bool)  {
	h.mutex.RLock()
	defer h.mutex.RUnlock()
	user, ok := h.hubs[id]
	return user, ok
}

func (h *DialogManager) Contains(id int64) bool {
	h.mutex.RLock()
	defer h.mutex.RUnlock()
	_, b := h.hubs[id]
	return b
}

func (h *DialogManager) Count() int {
	return len(h.hubs)
}
