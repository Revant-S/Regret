package client

import (
	"log"
	"sync"
)

var (
	once sync.Once
	tick *tickCounter
)

type tickCounter struct {
	tick uint64
	mu   sync.Mutex
}

func (t *tickCounter) Update(tick uint64) {
	t.mu.Lock()
	defer t.mu.Unlock()
	if tick <= t.tick {
		log.Fatalf("Update tick should be strictly more than current value")
	}
	t.tick = tick
}

func (t *tickCounter) Get() uint64 {
	return t.tick
}

func (t *tickCounter) Reset() {
	t.tick = 0
}

type TickCounterInterface interface {
	Update(tick uint64)
	Get() uint64
	Reset()
}

func GetTickInstance() TickCounterInterface {
	once.Do(func() {
		tick = &tickCounter{
			tick: 0,
		}
	})
	return tick
}
