package unique

type Ints64 struct {
	values map[int64]struct{}
}

func NewInts64(keys ...int64) Ints64 {
	k := Ints64{
		values: make(map[int64]struct{}),
	}
	for _, key := range keys {
		k.values[key] = struct{}{}
	}
	return k
}

func (k Ints64) Contains(val int64) bool {
	_, ok := k.values[val]

	return ok
}

func (k Ints64) Val() []int64 {
	keys := make([]int64, len(k.values))
	i := 0
	for key := range k.values {
		keys[i] = key
		i++
	}
	return keys
}

func (k Ints64) Len() int {
	return len(k.values)
}
