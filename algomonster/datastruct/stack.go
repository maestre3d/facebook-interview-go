package datastruct

// Stack
// Not safe for concurrent-scenarios
type Stack struct {
	items []interface{}
}

func (s *Stack) Push(i interface{}) {
	s.items = append(s.items, i)
}

func (s *Stack) Pop() interface{} {
	if s.IsEmpty() {
		return nil
	}

	var i interface{}
	i, s.items = s.items[len(s.items)-1], s.items[:len(s.items)-1]
	return i
}

func (s Stack) IsEmpty() bool {
	return len(s.items) == 0
}

func (s Stack) Len() int {
	return len(s.items)
}
