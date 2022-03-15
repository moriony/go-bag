package bag

import "errors"

var ErrNotFound = errors.New("item not found")

type Bag interface {
	Strings() Strings
	Bools() Bools
	Ints() Ints
	Ints64() Ints64
	Floats32() Floats32
	Floats64() Floats64
	StringSlices() StringSlices
	Unique() Unique
}

type bag struct {
	strings      Strings
	bools        Bools
	ints         Ints
	ints64       Ints64
	floats32     Floats32
	floats64     Floats64
	stringSlices StringSlices
	unique       Unique
}

func New() Bag {
	return &bag{
		strings:      NewStrings(),
		bools:        NewBools(),
		ints:         NewInts(),
		ints64:       NewInts64(),
		floats32:     NewFloats32(),
		floats64:     NewFloats64(),
		stringSlices: NewStringSlices(),
		unique:       NewUnique(),
	}
}

func (b *bag) Strings() Strings {
	return b.strings
}

func (b *bag) Bools() Bools {
	return b.bools
}

func (b *bag) Ints() Ints {
	return b.ints
}

func (b *bag) Ints64() Ints64 {
	return b.ints64
}

func (b *bag) Floats32() Floats32 {
	return b.floats32
}

func (b *bag) Floats64() Floats64 {
	return b.floats64
}

func (b *bag) StringSlices() StringSlices {
	return b.stringSlices
}

func (b *bag) Unique() Unique {
	return b.unique
}
