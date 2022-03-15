package bag

import "sync"

type StringSlices interface {
	Set(name string, value []string)
	Get(name string) ([]string, error)
}

func NewStringSlices() StringSlices {
	return &stringSlices{
		m:     sync.RWMutex{},
		items: make(map[string][]string),
	}
}

type stringSlices struct {
	m     sync.RWMutex
	items map[string][]string
}

func (b *stringSlices) Get(name string) ([]string, error) {
	b.m.RLock()
	defer b.m.RUnlock()

	val, ok := b.items[name]
	if !ok {
		return nil, ErrNotFound
	}

	result := make([]string, len(val))
	copy(result, val)

	return result, nil
}

func (b *stringSlices) Set(name string, value []string) {
	b.m.Lock()
	b.items[name] = make([]string, len(value))
	copy(b.items[name], value)
	b.m.Unlock()
}
