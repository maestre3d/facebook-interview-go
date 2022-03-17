package datastruct

type CircularQueue struct {
	items       []interface{}
	size        int
	headPointer int
	tailPointer int
}

func NewCircularQueue(s int) CircularQueue {
	return CircularQueue{
		items:       make([]interface{}, s),
		size:        s,
		headPointer: 0,
		tailPointer: 0,
	}
}

func (q *CircularQueue) Push(i interface{}) {
	q.tailPointer++
	q.tailPointer %= q.size
	q.items[q.tailPointer] = i
}

func (q *CircularQueue) Pop() interface{} {
	q.headPointer++
	q.headPointer %= q.size
	i := q.items[q.headPointer]
	return i
}
