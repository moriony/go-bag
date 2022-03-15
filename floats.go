package bag

import "sync"

type Floats32 interface {
	Set(name string, value float32)
	Get(name string) (float32, error)
}

func NewFloats32() Floats32 {
	return &floats32{
		m:     sync.RWMutex{},
		items: make(map[string]float32),
	}
}

type floats32 struct {
	m     sync.RWMutex
	items map[string]float32
}

func (b *floats32) Get(name string) (float32, error) {
	b.m.RLock()
	val, ok := b.items[name]
	b.m.RUnlock()

	if !ok {
		return 0, ErrNotFound
	}
	return val, nil
}

func (b *floats32) Set(name string, value float32) {
	b.m.Lock()
	b.items[name] = value
	b.m.Unlock()
}

type Floats64 interface {
	Set(name string, value float64)
	Get(name string) (float64, error)
}

func NewFloats64() Floats64 {
	return &floats64{
		m:     sync.RWMutex{},
		items: make(map[string]float64),
	}
}

type floats64 struct {
	m     sync.RWMutex
	items map[string]float64
}

func (b *floats64) Get(name string) (float64, error) {
	b.m.RLock()
	val, ok := b.items[name]
	b.m.RUnlock()

	if !ok {
		return 0, ErrNotFound
	}
	return val, nil
}

func (b *floats64) Set(name string, value float64) {
	b.m.Lock()
	b.items[name] = value
	b.m.Unlock()
}
