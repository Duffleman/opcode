package opcode

type MemoryStore struct {
	cursor       int
	RelativeBase int
	v            []int
}

func NewMemStore(start []int, defaultSize *int) *MemoryStore {
	if defaultSize == nil {
		return &MemoryStore{
			cursor: 0,
			v:      start,
		}
	}

	size := *defaultSize

	s := make([]int, size)

	for k, v := range start {
		s[k] = v
	}

	return &MemoryStore{
		cursor: 0,
		v:      s,
	}
}

func (ms *MemoryStore) Get(pos int) int {
	return ms.v[pos]
}

func (ms *MemoryStore) GetAt(pos, paramMode int) int {
	switch paramMode {
	case 0:
		return ms.Get(ms.Get(pos))
	case 1:
		return ms.Get(pos)
	case 2:
		return ms.Get(ms.Get(pos) + ms.RelativeBase)
	default:
		return 0
	}
}

func (ms *MemoryStore) Set(pos, val int) {
	ms.v[pos] = val
}

func (ms *MemoryStore) Next() (int, bool) {
	defer func() { ms.cursor++ }()

	if ms.cursor >= len(ms.v) {
		return ms.cursor, false
	}

	return ms.cursor, true
}

func (ms *MemoryStore) Jump(to int) {
	ms.cursor = to
}

func (ms *MemoryStore) All() []int {
	return ms.v
}
