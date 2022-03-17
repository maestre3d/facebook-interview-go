package datastruct

type Queue struct {
	items []interface{}
}

func (q Queue) Len() int {
	return len(q.items)
}

func (q Queue) IsEmpty() bool {
	return len(q.items) == 0
}

func (q *Queue) Push(i interface{}) {
	q.items = append(q.items, i)
}

func (q *Queue) Pop() interface{} {
	if q.IsEmpty() {
		return nil
	}
	i := q.items[0]
	if q.Len() == 1 {
		// no more items after pop, avoid overflow
		q.items = []interface{}{}
		return i
	}
	q.items = q.items[1:]
	return i
}
