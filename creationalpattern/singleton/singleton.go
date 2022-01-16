package singleton

import (
	"sync"
)

var (
	instance *counter
	once     sync.Once
)

type counter struct {
	number int
	mux    sync.RWMutex
}

func (c *counter) Add(n int) {
	c.mux.Lock()
	defer c.mux.Unlock()

	c.number += n
}

func (c *counter) Sub(n int) {
	c.mux.Lock()
	defer c.mux.Unlock()

	c.number -= n
}

func (c *counter) Get() int {
	c.mux.RLock()
	defer c.mux.RUnlock()

	return c.number
}

func GetInstance() *counter {
	once.Do(func() {
		instance = &counter{
			number: 0,
		}
	})

	return instance
}
