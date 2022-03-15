package bag

type Unique interface {
	Ints64() UniqueInts64
}

func NewUnique() Unique {
	return &uniq{
		ints64: NewUniqueInts64(),
	}
}

type uniq struct {
	ints64 UniqueInts64
}

func (u *uniq) Ints64() UniqueInts64 {
	return u.ints64
}
