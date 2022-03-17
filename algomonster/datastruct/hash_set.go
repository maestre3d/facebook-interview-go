package datastruct

// HashSet is useful when you only need to know existence of a key. Example use cases include DFS and BFS on graphs.
type HashSet map[interface{}]struct{}

func (s HashSet) Contains(v interface{}) bool {
	_, ok := s[v]
	return ok
}

func (s HashSet) Add(v interface{}) {
	s[v] = struct{}{}
}

func (s HashSet) Remove(v interface{}) {
	delete(s, v)
}
