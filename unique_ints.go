package bag

import (
	"sync"

	"github.com/moriony/go-bag/unique"
)

type UniqueInts64 interface {
	Set(name string, value unique.Ints64)
	Get(name string) (unique.Ints64, error)
}

func NewUniqueInts64() UniqueInts64 {
	return &uniqInts64{
		m:     sync.RWMutex{},
		items: make(map[string]unique.Ints64),
	}
}

type uniqInts64 struct {
	m     sync.RWMutex
	items map[string]unique.Ints64
}

func (b *uniqInts64) Get(name string) (unique.Ints64, error) {
	b.m.RLock()
	val, ok := b.items[name]
	b.m.RUnlock()

	if !ok {
		return unique.Ints64{}, ErrNotFound
	}
	return val, nil
}

func (b *uniqInts64) Set(name string, value unique.Ints64) {
	b.m.Lock()
	b.items[name] = value
	b.m.Unlock()
}
