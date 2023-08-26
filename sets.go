package mango

var val = struct{}{}

type Set[K comparable] map[K]struct{}

func NewSet[K comparable](items ...K) Set[K] {
	st := make(Set[K])
	for _, item := range items {
		st.Add(item)
	}
	return st
}

func (s Set[K]) Has(item K) bool {
	_, ok := s[item]
	return ok
}

func (s Set[K]) Add(item K) {
	s[item] = val
}

func (s Set[K]) Delete(item K) bool {
	if !s.Has(item) {
		return false
	}
	delete(s, item)
	return true
}

func (s Set[K]) Size() int {
	return len(s)
}
