package db

type M struct {
	MainMap    M1
	ReverseMap M2
}

// Stack of structs. Each struct has MainMap and ReverseMap.
type Stack []M

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(m1 M1, m2 M2) {
	*s = append(*s, M{m1, m2})
}

func (s *Stack) Pop() (M, bool) {
	if s.IsEmpty() {
		return M{}, false
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]

		return element, true
	}
}
