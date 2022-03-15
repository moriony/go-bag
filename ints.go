package bag

import "sync"

type Ints interface {
	Set(name string, value int)
	Get(name string) (int, error)
}

func NewInts() Ints {
	return &ints{
		m:     sync.RWMutex{},
		items: make(map[string]int),
	}
}

type ints struct {
	m     sync.RWMutex
	items map[string]int
}

func (b *ints) Get(name string) (int, error) {
	b.m.RLock()
	val, ok := b.items[name]
	b.m.RUnlock()

	if !ok {
		return 0, ErrNotFound
	}
	return val, nil
}

func (b *ints) Set(name string, value int) {
	b.m.Lock()
	b.items[name] = value
	b.m.Unlock()
}

type Ints64 interface {
	Set(name string, value int64)
	Get(name string) (int64, error)
}

func NewInts64() Ints64 {
	return &ints64{
		m:     sync.RWMutex{},
		items: make(map[string]int64),
	}
}

type ints64 struct {
	m     sync.RWMutex
	items map[string]int64
}

func (b *ints64) Get(name string) (int64, error) {
	b.m.RLock()
	val, ok := b.items[name]
	b.m.RUnlock()

	if !ok {
		return 0, ErrNotFound
	}
	return val, nil
}

func (b *ints64) Set(name string, value int64) {
	b.m.Lock()
	b.items[name] = value
	b.m.Unlock()
}
