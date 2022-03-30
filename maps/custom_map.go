package maps

import "sync"

func createRWLockMap(hint int) *customMap {
	return &customMap{
		a: make(map[any]any, hint),
	}
}

var _ Map = createRWLockMap(0)

type customMap struct {
	a    map[any]any
	lock sync.RWMutex
}

func (c *customMap) Set(k, v any) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.a[k] = v

}

func (c *customMap) Del(k any) {
	c.lock.Lock()
	defer c.lock.Unlock()
	delete(c.a, k)
}

func (c *customMap) Get(k any) (any, bool) {
	c.lock.RLock()
	defer c.lock.RUnlock()
	a, ok := c.a[k]
	return a, ok
}
