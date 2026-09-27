package clock

import (
	"sync"
)

type clock struct {
	tick uint64
	mu   sync.Mutex
}

var (
	clockInstance *clock
	once          sync.Once
)

type ClockInterface interface {
	UpdateTick(updateTo uint64)
	StartTick()
	GetTick() uint64
}

func GetClockInstance() ClockInterface {
	once.Do(func() {
		clockInstance = &clock{
			tick: 0,
		}
	})
	return clockInstance
}

func (c *clock) UpdateTick(updateTo uint64) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.tick = max(c.tick, updateTo)
}

func (c *clock) StartTick() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.tick = 0
}

func (c *clock) GetTick() uint64 {
	return c.tick
}
