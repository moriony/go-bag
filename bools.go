package bag

import "sync"

type Bools interface {
	Set(name string, value bool)
	Get(name string) (bool, error)
}

func NewBools() Bools {
	return &bools{
		m:     sync.RWMutex{},
		items: make(map[string]bool),
	}
}

type bools struct {
	m     sync.RWMutex
	items map[string]bool
}

func (b *bools) Get(name string) (bool, error) {
	b.m.RLock()
	val, ok := b.items[name]
	b.m.RUnlock()

	if !ok {
		return false, ErrNotFound
	}
	return val, nil
}

func (b *bools) Set(name string, value bool) {
	b.m.Lock()
	b.items[name] = value
	b.m.Unlock()
}
