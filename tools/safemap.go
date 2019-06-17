package tools

import (
	"sync"
	"time"
)

type safeobj struct {
	timeout int64
	obj     interface{}
}
type SafeMap struct {
	Data map[string]safeobj
	*sync.RWMutex
}

func (d SafeMap) Len() int {
	d.RLock()
	defer d.RUnlock()
	return len(d.Data)
}

func (d SafeMap) Contain(k string) bool {
	d.RLock()
	defer d.RUnlock()
	_, ok := d.Data[k]
	return ok
}

func (d SafeMap) Get(k string) interface{} {
	d.RLock()
	defer d.RUnlock()
	obj, ok := d.Data[k]
	if ok {
		if obj.timeout != -1 && time.Now().Unix() > obj.timeout {
			delete(d.Data, k)
			return nil
		}
		return obj.obj
	}

	return nil
}

func (d SafeMap) Set(k string, v interface{}) {
	d.Lock()
	defer d.Unlock()
	d.Data[k] = safeobj{-1, v}
}

func (d SafeMap) Set2(k string, v interface{}, timeout int64) {
	d.Lock()
	defer d.Unlock()
	d.Data[k] = safeobj{timeout, v}
}

func (d SafeMap) Remove(k string) {
	d.Lock()
	defer d.Unlock()
	delete(d.Data, k)
}

var safemap *SafeMap

func GetSafeMap() *SafeMap {
	if safemap != nil {
		return safemap
	}
	safemap = new(SafeMap)
	safemap.Data = map[string]safeobj{}
	safemap.RWMutex = &sync.RWMutex{}
	return safemap
}

func NewSafeMap() *SafeMap {
	sfm := new(SafeMap)
	sfm.Data = map[string]safeobj{}
	sfm.RWMutex = &sync.RWMutex{}
	return sfm
}
