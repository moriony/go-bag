package bag

import "sync"

type Strings interface {
	Set(name, value string)
	Get(name string) (string, error)
}

func NewStrings() Strings {
	return &strings{
		m:     sync.RWMutex{},
		items: make(map[string]string),
	}
}

type strings struct {
	m     sync.RWMutex
	items map[string]string
}

func (b *strings) Get(name string) (string, error) {
	b.m.RLock()
	val, ok := b.items[name]
	b.m.RUnlock()

	if !ok {
		return "", ErrNotFound
	}
	return val, nil
}

func (b *strings) Set(name, value string) {
	b.m.Lock()
	b.items[name] = value
	b.m.Unlock()
}
